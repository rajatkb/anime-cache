package services

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

const (
	numOperations = 100000000
	numConcurrent = 50
	maxKey        = 100000
)

func TestConcurrentOperationsWithThroughputReport(t *testing.T) {
	store := NewStore[int64](numConcurrent)

	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()

	for i := 0; i < numConcurrent; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				key := int64(rand.Intn(maxKey))
				value := key * 2
				store.Put(key, value)
			}
		}()
	}

	wg.Wait()
	opsCounter := numOperations * numConcurrent

	elapsed := time.Since(startTime)
	opsPerMillisecond := float64(opsCounter) / float64(elapsed.Milliseconds())
	fmt.Printf("Operations: %d, Elapsed Time: %s, Ops/ms: %.2f", opsCounter, elapsed, opsPerMillisecond)

}
