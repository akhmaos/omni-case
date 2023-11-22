package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/akhmaos/omni-case/internal/service"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
	processFunc   func(ctx context.Context, batch service.Batch) error
	getLimitsFunc func() (limit uint64, timeout time.Duration)
}

func (m *mockService) Process(ctx context.Context, batch service.Batch) error {
	if m.processFunc != nil {
		return m.processFunc(ctx, batch)
	}
	return nil
}

func (m *mockService) GetLimits() (limit uint64, timeout time.Duration) {
	if m.getLimitsFunc != nil {
		return m.getLimitsFunc()
	}
	return 0, 0
}

func TestProcessItems(t *testing.T) {
	mockService := &mockService{
		processFunc: func(ctx context.Context, batch service.Batch) error {
			return nil
		},
		getLimitsFunc: func() (limit uint64, timeout time.Duration) {
			return 2, time.Second
		},
	}

	client := NewClient(mockService)

	ctx := context.TODO()

	items := []service.Item{}

	err := client.ProcessItems(ctx, items)

	assert.NoError(t, err)
}

func TestProcessItemsWithError(t *testing.T) {
	mockService := &mockService{
		processFunc: func(ctx context.Context, batch service.Batch) error {
			return errors.New("test error")
		},
		getLimitsFunc: func() (limit uint64, timeout time.Duration) {
			return 2, time.Second
		},
	}

	client := NewClient(mockService)

	ctx := context.TODO()

	items := []service.Item{}

	err := client.ProcessItems(ctx, items)

	assert.Error(t, err)
}
