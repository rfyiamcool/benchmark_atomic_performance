# benchmark_atomic_performance

golang atomic performance

### performance

cmd

```
go test --bench=.
```

stdout (data: 2019-06)

```
goos: darwin
goarch: amd64
pkg: github.com/rfyiamcool/benchmark_atomic_performance
BenchmarkAtomicStore-4                 	200000000	         7.67 ns/op
BenchmarkNormalStore-4                 	2000000000	         0.42 ns/op
BenchmarkAtomicAdd-4                   	200000000	         7.69 ns/op
BenchmarkNormalAdd-4                   	2000000000	         1.67 ns/op
BenchmarkAtomicCas-4                   	200000000	         7.65 ns/op
BenchmarkNormalCas-4                   	2000000000	         0.64 ns/op
BenchmarkAtomicStoreParallel100-4      	100000000	        14.8 ns/op
BenchmarkAtomicCasShareParallel100-4   	100000000	        22.5 ns/op
BenchmarkAtomicCasParallel100-4        	300000000	         4.02 ns/op
BenchmarkAtomicAddParallel100-4        	100000000	        17.5 ns/op
BenchmarkAtomicAddParallel500-4        	100000000	        17.0 ns/op
```

conclusion

```
not parallel atomic ≈ 7.5ns
parallel atomic ≈ 15ns
```
