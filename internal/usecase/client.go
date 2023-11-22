package usecase

import (
	"context"
	"math"
	"time"

	"github.com/akhmaos/omni-case/internal"
	"github.com/akhmaos/omni-case/internal/service"
)

// Client - service for preparing items before send to external service
type Client interface {
	ProcessItems(ctx context.Context, items []service.Item) error
}

type client struct {
	service internal.Service
}

// NewClient - returning client implementation
func NewClient(service internal.Service) Client {
	return &client{
		service: service,
	}
}

func (c *client) ProcessItems(ctx context.Context, items []service.Item) error {
	limit, timeout := c.service.GetLimits()

	// проверка на необходимость разделения наших итемов на несколько батчей
	if int(limit) > len(items) {
		err := c.service.Process(ctx, service.Batch(items))
		if err != nil {
			return err
		}
		return nil
	}

	// Округление в большую сторону для создания бакетов под батчи
	batchCount := int(math.Ceil(float64(len(items))/float64(int(limit)))) - 1
	batchItemsCh := make(chan []service.Item, batchCount)

	errCh := make(chan error)

	go func(chan<- []service.Item) {
		defer func() {
			close(batchItemsCh)
		}()
		for i := 0; i < len(items); i++ {
			if i+int(limit) < len(items) {
				batchItemsCh <- items[i : i+int(limit)]
			} else {
				batchItemsCh <- items[i:]
			}
			i += int(limit) - 1
		}
	}(batchItemsCh)

	go func(<-chan []service.Item) {
		for itms := range batchItemsCh {
			err := c.service.Process(ctx, service.Batch(itms))
			if err != nil {
				errCh <- err
			}
			time.Sleep(timeout)
		}
	}(batchItemsCh)

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}
