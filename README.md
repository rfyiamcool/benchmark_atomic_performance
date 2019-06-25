# benchmark_atomic_performance

```
go test --bench=.
```

performance

```
goos: darwin
goarch: amd64
pkg: github.com/benchmark_atomic_performance
BenchmarkAtomicStore-4         	200000000	         7.33 ns/op
BenchmarkNormalStore-4         	2000000000	         0.44 ns/op
BenchmarkAtomicAdd-4           	200000000	         7.56 ns/op
BenchmarkNormalAdd-4           	2000000000	         1.66 ns/op
BenchmarkAtomicCas-4           	200000000	         7.34 ns/op
BenchmarkNormalCas-4           	2000000000	         0.66 ns/op
BenchmarkAtomicParallel-4      	100000000	        15.0 ns/op
BenchmarkAtomicAddParallel-4   	100000000	        16.0 ns/op
```