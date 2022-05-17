package code_examples

import "testing"

// fun fact: array is not always faster!
//BenchmarkSliceGlobal-8   	  926196	      1257.0 ns/op	       0 B/op	       0 allocs/op
//BenchmarkArrayGlobal-8   	 2110324	       567.0 ns/op	       0 B/op	       0 allocs/op
//BenchmarkSliceLocal-8    	 2275382	       535.0 ns/op	       0 B/op	       0 allocs/op
//BenchmarkArrayLocal-8    	 1802491	       647.4 ns/op	       0 B/op	       0 allocs/op

var gs = make([]byte, 1000) // Global slice
var ga [1000]byte           // Global array

func BenchmarkSliceGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j, v := range gs {
			gs[j]++
			gs[j] = gs[j] + v + 10
			gs[j] += v
		}
	}
}

func BenchmarkArrayGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j, v := range ga {
			ga[j]++
			ga[j] = ga[j] + v + 10
			ga[j] += v
		}
	}
}

func BenchmarkSliceLocal(b *testing.B) {
	var s = make([]byte, 1000)
	for i := 0; i < b.N; i++ {
		for j, v := range s {
			s[j]++
			s[j] = s[j] + v + 10
			s[j] += v
		}
	}
}

func BenchmarkArrayLocal(b *testing.B) {
	var a [1000]byte
	for i := 0; i < b.N; i++ {
		for j, v := range a {
			a[j]++
			a[j] = a[j] + v + 10
			a[j] += v
		}
	}
}
