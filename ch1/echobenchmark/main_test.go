// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package echo

import "testing"

var largs = []string{"arg0", "This", "is", "a", "benchmark", "test"}

func benchmarkEchoFunc(n int, f func([]string)) {
	for i := 0; i < n; i++ {
		f(largs)
	}
}

// BenchmarkEcho3 is a trivial benchmark with static data
func BenchmarkEcho1(b *testing.B) {
	benchmarkEchoFunc(b.N, echo2)
}

// BenchmarkEcho3 is a trivial benchmark with static data
func BenchmarkEcho2(b *testing.B) {
	benchmarkEchoFunc(b.N, echo2)
}

// BenchmarkEcho3 is a trivial benchmark with static data
func BenchmarkEcho3(b *testing.B) {
	benchmarkEchoFunc(b.N, echo3)
}
