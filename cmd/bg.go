package main

import (
	"sync"

	"github.com/rs/zerolog"
)

type backgroundRunner struct {
	wg     *sync.WaitGroup
	logger zerolog.Logger
}

func newBackgroundRunner(logger zerolog.Logger) *backgroundRunner {
	return &backgroundRunner{
		wg:     &sync.WaitGroup{},
		logger: logger,
	}
}

func (b *backgroundRunner) Run(fn func()) {
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

func (b *backgroundRunner) Wait() {
	b.wg.Wait()
}
