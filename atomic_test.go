package main

import (
	"runtime"
	"sync/atomic"
	"testing"
)

var a int64

func BenchmarkAtomicStore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomic.StoreInt64(&a, int64(i))
	}
}

func BenchmarkNormalStore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a = 1
	}
}

func BenchmarkAtomicAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&a, 1)
	}
}

func BenchmarkNormalAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a++
	}
}

func BenchmarkAtomicCas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomic.CompareAndSwapInt64(&a, 0, int64(i))
	}
}

func BenchmarkNormalCas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if a == 0 {
			a = int64(i)
		}
	}
}

func BenchmarkAtomicStoreParallel100(b *testing.B) {
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.StoreInt64(&a, int64(0))
		}
	})
}

func BenchmarkAtomicCasShareParallel100(b *testing.B) {
	b.SetParallelism(100)
	var bb int64 = 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.CompareAndSwapInt64(&bb, bb, bb+1)
		}
	})
}

func BenchmarkAtomicCasParallel100(b *testing.B) {
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		var bb int64 = 0
		for pb.Next() {
			atomic.CompareAndSwapInt64(&bb, bb, bb+1)
		}
	})
}

func BenchmarkAtomicAddParallel100(b *testing.B) {
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt64(&a, 1)
		}
	})
}

func BenchmarkAtomicAddParallel500(b *testing.B) {
	runtime.GOMAXPROCS(2 * runtime.NumCPU())
	b.SetParallelism(500)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt64(&a, 1)
		}
	})
}
