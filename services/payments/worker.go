package payments

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"

	client "github.com/brave-intl/bat-go/libs/clients"
	appctx "github.com/brave-intl/bat-go/libs/context"
	"github.com/brave-intl/bat-go/libs/httpsignature"
	"github.com/brave-intl/bat-go/libs/logging"
	"github.com/brave-intl/bat-go/libs/payments"
	"github.com/brave-intl/bat-go/libs/redisconsumer"
	"github.com/google/uuid"
)

// Worker for payments
type Worker struct {
	rc redisconsumer.StreamClient
}

// NewWorker from redis client
func NewWorker(rc redisconsumer.StreamClient) *Worker {
	return &Worker{rc}
}

// HandlePrepareMessage by sending it to the payments service
func (w *Worker) HandlePrepareMessage(ctx context.Context, stream, id string, data []byte) error {
	client, err := client.New("https://nitro-payments.bsg.brave.software", "")
	if err != nil {
		return err
	}
	return w.requestHandler(ctx, client, "POST", "/v1/prepare", stream, id, data)
}

// HandleSubmitMessage by sending it to the payments service
func (w *Worker) HandleSubmitMessage(ctx context.Context, stream, id string, data []byte) error {
	client, err := client.New("https://nitro-payments.bsg.brave.software", "")
	if err != nil {
		return err
	}
	return w.requestHandler(ctx, client, "POST", "/v1/submit", stream, id, data)
}

// requestHandler is a generic handler for sending encapsulated http requests and storing the results
func (w *Worker) requestHandler(ctx context.Context, client *client.SimpleHTTPClient, method, path string, stream, id string, data []byte) error {
	logger, err := appctx.GetLogger(ctx)
	if err != nil {
		return err
	}

	isRetryBlocked, err := w.rc.GetMessageRetryAfter(ctx, id)
	if err != nil {
		return err
	}
	if isRetryBlocked {
		return errors.New("waiting for retry-after")
	}

	reqWrapper := payments.RequestWrapper{}
	err = json.Unmarshal(data, &reqWrapper)
	if err != nil {
		return err
	}

	r, err := client.NewRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	_, err = reqWrapper.Request.Extract(r)
	if err != nil {
		return err
	}

	// FIXME we should probably complete override the url based on params
	r.URL = client.BaseURL.ResolveReference(&url.URL{
		Path: r.URL.RequestURI(),
	})

	delay := 5 * time.Second
	resp, err := client.Do(ctx, r, nil)
	if resp != nil {
		retry := resp.Header.Get("x-retry-after")
		if retry != "" {
			tmp, err := strconv.Atoi(retry)
			if err != nil {
				logger.Error().Err(err).Msg("failed to parse x-retry-after header")
			}
			delay = time.Duration(tmp) * time.Second
		}
	}

	if err := w.rc.SetMessageRetryAfter(ctx, id, delay); err != nil {
		logger.Error().Err(err).Msg("failed to set retry-after key")
	}

	if err != nil {
		return err
	}
	if resp == nil {
		return errors.New("response was nil")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("response was not 200 OK")
	}

	sr, err := httpsignature.EncapsulateResponse(ctx, resp)
	if err != nil {
		return err
	}

	respWrapper := &payments.ResponseWrapper{
		ID:        uuid.New(),
		Timestamp: time.Now(),
		Response:  sr,
	}

	return w.rc.AddMessages(ctx, stream+payments.ResponseSuffix, respWrapper)
}

// HandlePrepareConfigMessage creates a new prepare consumer, waiting for all messages to be consumed
func (w *Worker) HandlePrepareConfigMessage(ctx context.Context, stream, id string, data []byte) error {
	return w.handleConfigMessage(w.HandlePrepareMessage, ctx, id, data)
}

// HandleSubmitConfigMessage creates a new submit consumer, waiting for all messages to be consumed
func (w *Worker) HandleSubmitConfigMessage(ctx context.Context, stream, id string, data []byte) error {
	return w.handleConfigMessage(w.HandleSubmitMessage, ctx, id, data)
}

// handleConfigMessage is a generic handler which creates a consumer, waiting for all messages to be consumed
func (w *Worker) handleConfigMessage(handle redisconsumer.MessageHandler, ctx context.Context, id string, data []byte) error {
	logger, err := appctx.GetLogger(ctx)
	if err != nil {
		return err
	}

	consumerCtx, cancelFunc := context.WithCancel(ctx)

	config := payments.WorkerConfig{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	ctx, logger = logging.UpdateContext(ctx, logger.With().Str("childGroup", config.ConsumerGroup).Logger())

	logger.Info().Msg("processed config")
	go func() {
		redisconsumer.StartConsumer(consumerCtx, w.rc, config.Stream, config.ConsumerGroup, "0", handle)
	}()

	for {
		lag, pending, err := w.rc.UnacknowledgedCounts(ctx, config.Stream, config.ConsumerGroup)
		if err != nil {
			logger.Error().Err(err).Msg("failed to get unacknowledged count")
		}
		if lag+pending == 0 {
			break
		}
		logger.Info().Int64("lag", lag).Int64("pending", pending).Msg("waiting")

		time.Sleep(10 * time.Second)
	}

	logger.Info().Msg("all messages handled")
	cancelFunc()

	return nil
}

// StartPrepareConfigConsumer is a convenience function for starting the prepare config consumer
func (w *Worker) StartPrepareConfigConsumer(ctx context.Context) error {
	return redisconsumer.StartConsumer(ctx, w.rc, payments.PrepareConfigStream, payments.PrepareConfigConsumerGroup, "0", w.HandlePrepareConfigMessage)
}

// StartSubmitConfigConsumer is a convenience function for starting the prepare config consumer
func (w *Worker) StartSubmitConfigConsumer(ctx context.Context) error {
	return redisconsumer.StartConsumer(ctx, w.rc, payments.SubmitConfigStream, payments.SubmitConfigConsumerGroup, "0", w.HandleSubmitConfigMessage)
}