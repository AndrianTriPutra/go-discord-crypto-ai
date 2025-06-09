package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func (a *apps) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Prepare signal handling
	stoped := make(chan os.Signal, 1)
	signal.Notify(
		stoped,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// Error channel and wait group
	errChan := make(chan error, 2) // buffer adjusted number of goroutines
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// Goroutine: Chatbot
	go func() {
		defer wg.Done()
		if err := a.chatbot(ctx); err != nil {
			errChan <- fmt.Errorf("[chatbot] failed: %w", err)
			cancel()
		}
	}()

	// Goroutine: Sniper
	go func() {
		defer wg.Done()
		if err := a.sniper(ctx); err != nil {
			errChan <- fmt.Errorf("[sniper] failed: %w", err)
			cancel()
		}
	}()

	// Select for stop signal or error
	var errN error
	select {
	case s := <-stoped:
		switch s {
		case syscall.SIGHUP:
			errN = errors.New("[hungup]")
		case syscall.SIGINT:
			errN = errors.New("[interrupt]")
		case syscall.SIGTERM:
			errN = errors.New("[force stop]")
		case syscall.SIGQUIT:
			errN = errors.New("[stop and core dump]")
		default:
			errN = errors.New("[unknown signal]")
		}
		cancel()

	case err := <-errChan:
		errN = err
		cancel()
	}

	wg.Wait()
	return errN
}
