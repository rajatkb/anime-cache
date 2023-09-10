package services

import (
	"animecache/entities"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func BenchmarkWriteOperations(t *testing.B) {

	type DeeplyNestedLargeStruct struct {
		Field1 int
		Field2 string
		Field3 struct {
			Nested1 int
			Nested2 string
			Nested3 struct {
				Deeper1 float64
				Deeper2 bool
				Deeper3 []int
			}
		}
		Field4 []struct {
			Item1 int
			Item2 string
		}
	}

	large := DeeplyNestedLargeStruct{
		Field1: 42,
		Field2: "Hello, World!",
		Field3: struct {
			Nested1 int
			Nested2 string
			Nested3 struct {
				Deeper1 float64
				Deeper2 bool
				Deeper3 []int
			}
		}{
			Nested1: 10,
			Nested2: "Nested Field",
			Nested3: struct {
				Deeper1 float64
				Deeper2 bool
				Deeper3 []int
			}{
				Deeper1: 3.14159265359,
				Deeper2: true,
				Deeper3: []int{1, 2, 3, 4, 5},
			},
		},
		Field4: []struct {
			Item1 int
			Item2 string
		}{
			{Item1: 1, Item2: "A"},
			{Item1: 2, Item2: "B"},
			{Item1: 3, Item2: "C"},
		},
	}

	const (
		numOperations = 2000
		numConcurrent = 2500
		maxKey        = 100000
	)
	storeTestList := []struct {
		storeName string
		store     entities.Store[*DeeplyNestedLargeStruct]
		runner    func(wg *sync.WaitGroup, store entities.Store[*DeeplyNestedLargeStruct])
	}{
		{
			storeName: "LockedMapStore",
			store:     NewLockedMapStore[*DeeplyNestedLargeStruct](),
			runner: func(wg *sync.WaitGroup, store entities.Store[*DeeplyNestedLargeStruct]) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := int64(rand.Intn(maxKey))
					store.Put(key, &large)
				}
			},
		},
		{
			storeName: "HaxMapStore",
			store:     NewHaxMapStore[*DeeplyNestedLargeStruct](numOperations),
			runner: func(wg *sync.WaitGroup, store entities.Store[*DeeplyNestedLargeStruct]) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := int64(rand.Intn(maxKey))
					store.Put(key, &large)
				}
			},
		},
	}

	storeTestList2 := []struct {
		storeName string
		store     entities.Store[DeeplyNestedLargeStruct]
		runner    func(wg *sync.WaitGroup, store entities.Store[DeeplyNestedLargeStruct])
	}{
		{
			storeName: "LockedMapStore",
			store:     NewLockedMapStore[DeeplyNestedLargeStruct](),
			runner: func(wg *sync.WaitGroup, store entities.Store[DeeplyNestedLargeStruct]) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := int64(rand.Intn(maxKey))
					store.Put(key, large)
				}
			},
		},
		{
			storeName: "HaxMapStore",
			store:     NewHaxMapStore[DeeplyNestedLargeStruct](numOperations),
			runner: func(wg *sync.WaitGroup, store entities.Store[DeeplyNestedLargeStruct]) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := int64(rand.Intn(maxKey))
					store.Put(key, large)
				}
			},
		},
	}

	for _, config := range storeTestList {

		t.Run(config.storeName+"_pointer", func(b *testing.B) {
			store := config.store

			var wg sync.WaitGroup
			rand.Seed(time.Now().UnixNano())

			startTime := time.Now()
			wg.Add(numConcurrent)

			b.ResetTimer()

			for i := 0; i < numConcurrent; i++ {
				go config.runner(&wg, store)
			}

			wg.Wait()
			opsCounter := numOperations * numConcurrent

			elapsed := time.Since(startTime)
			ops := float64(elapsed.Nanoseconds()) / float64(opsCounter)
			// fmt.Printf("Operations: %d, Elapsed Time: %s, Ops/Sec: %.2f \n", opsCounter, elapsed, ops)
			b.ReportMetric(ops, "ns/op")
			b.ReportMetric(10e9/ops, "ops/sec")
		})

	}

	for _, config := range storeTestList2 {

		t.Run(config.storeName+"_struct", func(b *testing.B) {
			store := config.store

			var wg sync.WaitGroup
			rand.Seed(time.Now().UnixNano())

			startTime := time.Now()
			wg.Add(numConcurrent)

			b.ResetTimer()

			for i := 0; i < numConcurrent; i++ {
				go config.runner(&wg, store)
			}

			wg.Wait()
			opsCounter := numOperations * numConcurrent

			elapsed := time.Since(startTime)
			ops := float64(elapsed.Nanoseconds()) / float64(opsCounter)
			// fmt.Printf("Operations: %d, Elapsed Time: %s, Ops/Sec: %.2f \n", opsCounter, elapsed, ops)
			b.ReportMetric(ops, "ns/op")
			b.ReportMetric(10e9/ops, "ops/sec")
		})

	}

}

func BenchmarkReadWriteOperations(t *testing.B) {

	type DeeplyNestedLargeStruct struct {
		Field1 int
		Field2 string
		Field3 struct {
			Nested1 int
			Nested2 string
			Nested3 struct {
				Deeper1 float64
				Deeper2 bool
				Deeper3 []int
			}
		}
		Field4 []struct {
			Item1 int
			Item2 string
		}
	}

	large := DeeplyNestedLargeStruct{
		Field1: 42,
		Field2: "Hello, World!",
		Field3: struct {
			Nested1 int
			Nested2 string
			Nested3 struct {
				Deeper1 float64
				Deeper2 bool
				Deeper3 []int
			}
		}{
			Nested1: 10,
			Nested2: "Nested Field",
			Nested3: struct {
				Deeper1 float64
				Deeper2 bool
				Deeper3 []int
			}{
				Deeper1: 3.14159265359,
				Deeper2: true,
				Deeper3: []int{1, 2, 3, 4, 5},
			},
		},
		Field4: []struct {
			Item1 int
			Item2 string
		}{
			{Item1: 1, Item2: "A"},
			{Item1: 2, Item2: "B"},
			{Item1: 3, Item2: "C"},
		},
	}

	const (
		numOperations = 2000
		numConcurrent = 2500
		maxKey        = 100000
	)

	storeTestList := []struct {
		storeName string
		store     entities.Store[DeeplyNestedLargeStruct]
		reader    func(wg *sync.WaitGroup, store entities.Store[DeeplyNestedLargeStruct])
		writer    func(wg *sync.WaitGroup, store entities.Store[DeeplyNestedLargeStruct])
	}{
		{
			storeName: "LockedMapStore",
			store:     NewLockedMapStore[DeeplyNestedLargeStruct](),
			reader: func(wg *sync.WaitGroup, store entities.Store[DeeplyNestedLargeStruct]) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := int64(rand.Intn(maxKey))
					store.Put(key, large)
				}
			},
			writer: func(wg *sync.WaitGroup, store entities.Store[DeeplyNestedLargeStruct]) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := int64(rand.Intn(maxKey))
					store.Get(key)
				}
			},
		},
		{
			storeName: "HaxMapStore",
			store:     NewHaxMapStore[DeeplyNestedLargeStruct](numOperations),
			reader: func(wg *sync.WaitGroup, store entities.Store[DeeplyNestedLargeStruct]) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := int64(rand.Intn(maxKey))
					store.Put(key, large)
				}
			},
			writer: func(wg *sync.WaitGroup, store entities.Store[DeeplyNestedLargeStruct]) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := int64(rand.Intn(maxKey))
					store.Get(key)
				}
			},
		},
	}

	storeTestList2 := []struct {
		storeName string
		store     entities.Store[*DeeplyNestedLargeStruct]
		reader    func(wg *sync.WaitGroup, store entities.Store[*DeeplyNestedLargeStruct])
		writer    func(wg *sync.WaitGroup, store entities.Store[*DeeplyNestedLargeStruct])
	}{
		{
			storeName: "LockedMapStore",
			store:     NewLockedMapStore[*DeeplyNestedLargeStruct](),
			reader: func(wg *sync.WaitGroup, store entities.Store[*DeeplyNestedLargeStruct]) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := int64(rand.Intn(maxKey))
					store.Put(key, &large)
				}
			},
			writer: func(wg *sync.WaitGroup, store entities.Store[*DeeplyNestedLargeStruct]) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := int64(rand.Intn(maxKey))
					store.Get(key)
				}
			},
		},
		{
			storeName: "HaxMapStore",
			store:     NewHaxMapStore[*DeeplyNestedLargeStruct](numOperations),
			reader: func(wg *sync.WaitGroup, store entities.Store[*DeeplyNestedLargeStruct]) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := int64(rand.Intn(maxKey))
					store.Put(key, &large)
				}
			},
			writer: func(wg *sync.WaitGroup, store entities.Store[*DeeplyNestedLargeStruct]) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := int64(rand.Intn(maxKey))
					store.Get(key)
				}
			},
		},
	}

	for _, config := range storeTestList {

		t.Run(config.storeName+"_struct", func(b *testing.B) {

			store := config.store

			var wg sync.WaitGroup
			rand.Seed(time.Now().UnixNano())

			startTime := time.Now()

			b.ResetTimer()

			for i := 0; i < numConcurrent; i++ {
				wg.Add(2)
				go config.writer(&wg, store)
				go config.reader(&wg, store)

			}

			wg.Wait()
			opsCounter := numOperations * numConcurrent

			elapsed := time.Since(startTime)
			ops := float64(elapsed.Nanoseconds()) / float64(opsCounter)
			// fmt.Printf("Operations: %d, Elapsed Time: %s, Ops/Sec: %.2f \n", opsCounter, elapsed, ops)

			b.ReportMetric(ops, "ns/op")
			b.ReportMetric(10e9/ops, "ops/sec")
		})

	}

	for _, config := range storeTestList2 {

		t.Run(config.storeName+"_pointer", func(b *testing.B) {

			store := config.store

			var wg sync.WaitGroup
			rand.Seed(time.Now().UnixNano())

			startTime := time.Now()

			b.ResetTimer()

			for i := 0; i < numConcurrent; i++ {
				wg.Add(2)
				go config.writer(&wg, store)
				go config.reader(&wg, store)

			}

			wg.Wait()
			opsCounter := numOperations * numConcurrent

			elapsed := time.Since(startTime)
			ops := float64(elapsed.Nanoseconds()) / float64(opsCounter)
			// fmt.Printf("Operations: %d, Elapsed Time: %s, Ops/Sec: %.2f \n", opsCounter, elapsed, ops)

			b.ReportMetric(ops, "ns/op")
			b.ReportMetric(10e9/ops, "ops/sec")
		})

	}

}
