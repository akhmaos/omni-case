package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Batch is a batch of items.
type Batch []Item

// Item is some abstract item.
type Item struct {
	Title string
	Key   string
}

// ErrBlocked - error when service is blocked
var ErrBlocked = errors.New("blocked")

// ExtServiceImpl is a simple implementation of the Service interface.
type ExtServiceImpl struct {
	mu           sync.Mutex
	maxItems     uint64
	timeout      time.Duration
	processing   uint64
	processingMu sync.Mutex
}

// NewService creates a new instance of ServiceImpl.
func NewService(maxItems uint64) *ExtServiceImpl {
	return &ExtServiceImpl{
		maxItems: maxItems,
		timeout:  5 * time.Second,
	}
}

// GetLimits returns the maximum number of items and processing time interval.
func (s *ExtServiceImpl) GetLimits() (uint64, time.Duration) {
	return s.maxItems - s.processing, s.timeout
}

// Process simulates the processing of a batch of items.
func (s *ExtServiceImpl) Process(ctx context.Context, batch Batch) error {

	// Check if the service is blocked.

	if s.processing > s.maxItems {
		go func() {
			s.processingMu.Lock()
			time.Sleep(25 * time.Second)
			defer s.processingMu.Unlock()
		}()
		return ErrBlocked
	}

	// Process items.
	s.processing += uint64(len(batch))
	fmt.Printf("Processed %d items\n", len(batch))

	go func() {
		time.Sleep(s.timeout)
		s.processing = 0
	}()
	return nil
}
