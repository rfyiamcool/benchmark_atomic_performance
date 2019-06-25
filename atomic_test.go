package main

import (
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

func BenchmarkAtomicParallel(b *testing.B) {
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.StoreInt64(&a, int64(0))
		}
	})
}

func BenchmarkAtomicAddParallel(b *testing.B) {
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt64(&a, 1)
		}
	})
}
