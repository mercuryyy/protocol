// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webhook

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"sync"
	"time"

	"github.com/frostbyte73/core"
	"github.com/hashicorp/go-retryablehttp"
	"go.uber.org/atomic"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
)

const (
	numWorkers       = 10
	defaultQueueSize = 100
)

type URLNotifierParams struct {
	HTTPClientParams
	Logger     logger.Logger
	QueueSize  int
	URL        string
	APIKey     string
	APISecret  string
	FieldsHook func(whi *livekit.WebhookInfo)
	FilterParams
}

// URLNotifier is a QueuedNotifier that sends a POST request to a Webhook URL.
// It will retry on failure, and will drop events if notification fall too far behind
type URLNotifier struct {
	mu            sync.RWMutex
	params        URLNotifierParams
	client        *retryablehttp.Client
	dropped       atomic.Int32
	pool          core.QueuePool
	processedHook func(ctx context.Context, whi *livekit.WebhookInfo)
	filter        *filter
}

func NewURLNotifier(params URLNotifierParams) *URLNotifier {
	if params.QueueSize == 0 {
		params.QueueSize = defaultQueueSize
	}
	if params.Logger == nil {
		params.Logger = logger.GetLogger()
	}

	rhc := retryablehttp.NewClient()
	if params.RetryWaitMin > 0 {
		rhc.RetryWaitMin = params.RetryWaitMin
	}
	if params.RetryWaitMax > 0 {
		rhc.RetryWaitMax = params.RetryWaitMax
	}
	if params.MaxRetries > 0 {
		rhc.RetryMax = params.MaxRetries
	}
	if params.ClientTimeout > 0 {
		rhc.HTTPClient.Timeout = params.ClientTimeout
	}
	n := &URLNotifier{
		params: params,
		client: rhc,
		filter: newFilter(params.FilterParams),
	}
	n.client.Logger = &logAdapter{}

	n.pool = core.NewQueuePool(numWorkers, core.QueueWorkerParams{
		QueueSize:    params.QueueSize,
		DropWhenFull: true,
	})
	return n
}

func (n *URLNotifier) SetKeys(apiKey, apiSecret string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.params.APIKey = apiKey
	n.params.APISecret = apiSecret
}

func (n *URLNotifier) SetFilter(params FilterParams) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.filter.SetFilter(params)
}

func (n *URLNotifier) RegisterProcessedHook(hook func(ctx context.Context, whi *livekit.WebhookInfo)) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.processedHook = hook
}

func (n *URLNotifier) getProcessedHook() func(ctx context.Context, whi *livekit.WebhookInfo) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.processedHook
}

func (n *URLNotifier) QueueNotify(ctx context.Context, event *livekit.WebhookEvent) error {
	if !n.filter.IsAllowed(event.Event) {
		return nil
	}

	enqueuedAt := time.Now()

	key := eventKey(event)
	if !n.pool.Submit(key, func() {
		fields := logFields(event, n.params.URL)

		queueDuration := time.Since(enqueuedAt)
		fields = append(fields, "queueDuration", queueDuration)

		sendStart := time.Now()
		err := n.send(event)
		sendDuration := time.Since(sendStart)
		fields = append(fields, "sendDuration", sendDuration)
		if err != nil {
			n.params.Logger.Warnw("failed to send webhook", err, fields...)
			n.dropped.Add(event.NumDropped + 1)
		} else {
			n.params.Logger.Infow("sent webhook", fields...)
		}
		if ph := n.getProcessedHook(); ph != nil {
			whi := webhookInfo(
				event,
				enqueuedAt,
				queueDuration,
				sendStart,
				sendDuration,
				n.params.URL,
				false,
				err,
			)
			if n.params.FieldsHook != nil {
				n.params.FieldsHook(whi)
			}
			ph(ctx, whi)
		}
	}) {
		n.dropped.Inc()

		fields := logFields(event, n.params.URL)
		n.params.Logger.Infow("dropped webhook", fields...)

		if ph := n.getProcessedHook(); ph != nil {
			whi := webhookInfo(
				event,
				time.Time{},
				0,
				time.Time{},
				0,
				n.params.URL,
				true,
				nil,
			)
			if n.params.FieldsHook != nil {
				n.params.FieldsHook(whi)
			}
			ph(ctx, whi)
		}
	}
	return nil
}

func (n *URLNotifier) Stop(force bool) {
	if force {
		n.pool.Kill()
	} else {
		n.pool.Drain()
	}
}

func (n *URLNotifier) send(event *livekit.WebhookEvent) error {
	// set dropped count
	event.NumDropped = n.dropped.Swap(0)
	encoded, err := protojson.Marshal(event)
	if err != nil {
		return err
	}
	// sign payload
	sum := sha256.Sum256(encoded)
	b64 := base64.StdEncoding.EncodeToString(sum[:])

	n.mu.RLock()
	apiKey := n.params.APIKey
	apiSecret := n.params.APISecret
	n.mu.RUnlock()

	at := auth.NewAccessToken(apiKey, apiSecret).
		SetValidFor(5 * time.Minute).
		SetSha256(b64)
	token, err := at.ToJWT()
	if err != nil {
		return err
	}
	r, err := retryablehttp.NewRequest("POST", n.params.URL, bytes.NewReader(encoded))
	if err != nil {
		// ignore and continue
		return err
	}
	r.Header.Set(authHeader, token)
	// use a custom mime type to ensure signature is checked prior to parsing
	r.Header.Set("content-type", "application/webhook+json")
	res, err := n.client.Do(r)
	if err != nil {
		return err
	}
	_ = res.Body.Close()
	return nil
}
