package internal

import (
	"context"
	"time"

	"github.com/akhmaos/omni-case/internal/service"
)

// Service defines external service that can process batches of items.
type Service interface {
	GetLimits() (n uint64, p time.Duration)
	Process(ctx context.Context, batch service.Batch) error
}
