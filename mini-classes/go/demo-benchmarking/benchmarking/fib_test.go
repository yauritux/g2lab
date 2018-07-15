package benchmarking

import "testing"

func BenchmarkFib20(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		//Fib(20) // run the Fib function b.N times
		FibNew(20) // run the FibNew function b.N times
	}
}
