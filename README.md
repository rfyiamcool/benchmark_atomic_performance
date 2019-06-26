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
BenchmarkAtomicStore-4                 	200000000	         7.56 ns/op
BenchmarkNormalStore-4                 	2000000000	         0.40 ns/op
BenchmarkAtomicAdd-4                   	200000000	         7.77 ns/op
BenchmarkNormalAdd-4                   	2000000000	         1.64 ns/op
BenchmarkAtomicCas-4                   	200000000	         7.95 ns/op
BenchmarkNormalCas-4                   	2000000000	         0.74 ns/op
BenchmarkAtomicParallel100-4           	100000000	        13.1 ns/op
BenchmarkAtomicCasShareParallel100-4   	100000000	        22.7 ns/op
BenchmarkAtomicCasParallel100-4        	500000000	         3.48 ns/op
BenchmarkAtomicAddParallel100-4        	100000000	        14.4 ns/op
BenchmarkAtomicAddParallel500-4        	100000000	        14.4 ns/op
```

conclusion

```
not parallel atomic ≈ 7.5ns
parallel atomic ≈ 15ns
```
