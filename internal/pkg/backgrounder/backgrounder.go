package backgrounder

import (
	"sync"
	"time"
)

// BackgroundWorker manages the execution of functions at specific intervals.
type Worker struct {
	wg        *sync.WaitGroup
	mu        sync.Mutex
	isRunning map[string]bool
}

// NewBackgroundWorker creates a new instance of BackgroundWorker.
func NewBackgroundWorker(wg *sync.WaitGroup) *Worker {
	return &Worker{
		wg:        wg,
		isRunning: make(map[string]bool),
	}
}

// RunWithInterval starts a background thread to run the given function at the specified interval (in milliseconds).
// If the function is still running, it waits for the current execution to finish before scheduling the next.
func (bw *Worker) RunWithInterval(name string, fn func(), intervalMs int) {
	bw.wg.Add(1)
	go func() {
		defer bw.wg.Done()
		// Run the function immediately
		bw.mu.Lock()
		if !bw.isRunning[name] {
			bw.isRunning[name] = true
			bw.mu.Unlock()

			fn()

			bw.mu.Lock()
			bw.isRunning[name] = false
			bw.mu.Unlock()
		} else {
			bw.mu.Unlock()
		}

		// Run the function at intervals
		ticker := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			bw.mu.Lock()
			if bw.isRunning[name] {
				bw.mu.Unlock()
				continue // Skip if the function is still running
			}
			bw.isRunning[name] = true
			bw.mu.Unlock()

			fn()

			bw.mu.Lock()
			bw.isRunning[name] = false
			bw.mu.Unlock()
		}
	}()
}
