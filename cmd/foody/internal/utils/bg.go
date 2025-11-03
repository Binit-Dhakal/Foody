package utils

import (
	"sync"

	"github.com/rs/zerolog"
)

type BackgroundRunner struct {
	wg     *sync.WaitGroup
	logger zerolog.Logger
}

func NewBackgroundRunner(logger zerolog.Logger) *BackgroundRunner {
	return &BackgroundRunner{
		wg:     &sync.WaitGroup{},
		logger: logger,
	}
}

func (b *BackgroundRunner) Run(fn func()) {
	b.wg.Add(1)
	go func() {
		defer b.wg.Done()
		defer func() {
			if err := recover(); err != nil {
				b.logger.Error().Err(err.(error)).Msg("failed background goroutine")
			}
		}()
		fn()
	}()
}

func (b *BackgroundRunner) Wait() {
	b.wg.Wait()
}
