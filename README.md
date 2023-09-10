# anime-cache
anime-cache or Another In Memory Cache is golang in memory cache intended for usage within the application runtime



## Benchmarks

Compares the following
- Map with RWLock
- [HaxMap](https://github.com/alphadose/haxmap/tree/main)
```
goos: linux
goarch: amd64
pkg: animecache/services
cpu: AMD Ryzen 5 5600X 6-Core Processor
BenchmarkWriteOperations
BenchmarkWriteOperations/LockedMapStore_pointer
BenchmarkWriteOperations/LockedMapStore_pointer-4               1000000000             107.2 ns/op        93254418 ops/sec             0 B/op          0 allocs/op
BenchmarkWriteOperations/HaxMapStore_pointer
BenchmarkWriteOperations/HaxMapStore_pointer-4                  1000000000             123.2 ns/op        81154940 ops/sec             0 B/op          0 allocs/op
BenchmarkWriteOperations/LockedMapStore_struct
BenchmarkWriteOperations/LockedMapStore_struct-4                       4               217.7 ns/op        45937588 ops/sec         84888 B/op       1196 allocs/op
BenchmarkWriteOperations/HaxMapStore_struct
BenchmarkWriteOperations/HaxMapStore_struct-4                          1               203.7 ns/op        49084570 ops/sec      573474848 B/op   5104766 allocs/op
BenchmarkReadWriteOperations
BenchmarkReadWriteOperations/LockedMapStore_struct
BenchmarkReadWriteOperations/LockedMapStore_struct-4                   1               371.9 ns/op        26890025 ops/sec      39826744 B/op      16184 allocs/op
BenchmarkReadWriteOperations/HaxMapStore_struct
BenchmarkReadWriteOperations/HaxMapStore_struct-4                      1               317.9 ns/op        31456883 ops/sec      569707632 B/op   5109499 allocs/op
BenchmarkReadWriteOperations/LockedMapStore_pointer
BenchmarkReadWriteOperations/LockedMapStore_pointer-4                  1               345.9 ns/op        28906078 ops/sec       6297384 B/op      13589 allocs/op
BenchmarkReadWriteOperations/HaxMapStore_pointer
BenchmarkReadWriteOperations/HaxMapStore_pointer-4                     1               303.2 ns/op        32984633 ops/sec      49654416 B/op    5109739 allocs/op
```