package cacheInMemoryMutexComparation

import (
	"cacheInMemoryMutexComparation/channel"
	"cacheInMemoryMutexComparation/mux"
	"testing"
)

/*
goos: linux
goarch: 386
pkg: cacheInMemoryMutexComparison
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
Benchmark_Mutex
Benchmark_Mutex-8     	    9747	    122358 ns/op
Benchmark_Channel
Benchmark_Channel-8   	    5052	    289512 ns/op
*/

func Benchmark_Mutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test := mux.NewCache()
		for i := 0; i < 100; i++ {
			go func() {
				test.Set("key", 1)
			}()
		}
		for i := 0; i < 100; i++ {
			go func() {
				test.Get("key")
			}()
		}
		for i := 0; i < 100; i++ {
			go func() {
				test.Delete("key")
			}()
		}
	}
}

func Benchmark_Channel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test := channel.NewCache()
		for i := 0; i < 100; i++ {
			go func() {
				test.Set("key", 1)
			}()
		}
		for i := 0; i < 100; i++ {
			go func() {
				test.Get("key")
			}()
		}
		for i := 0; i < 100; i++ {
			go func() {
				test.Delete("key")
			}()
		}
	}
}
