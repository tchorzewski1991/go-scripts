package fib

import (
	"testing"
)

func BenchmarkRecursiveFib1(b *testing.B)  { benchmarkRecursiveFibFor(1, b) }
func BenchmarkRecursiveFib2(b *testing.B)  { benchmarkRecursiveFibFor(2, b) }
func BenchmarkRecursiveFib5(b *testing.B)  { benchmarkRecursiveFibFor(5, b) }
func BenchmarkRecursiveFib10(b *testing.B) { benchmarkRecursiveFibFor(10, b) }
func BenchmarkRecursiveFib20(b *testing.B) { benchmarkRecursiveFibFor(20, b) }
func BenchmarkRecursiveFib40(b *testing.B) { benchmarkRecursiveFibFor(40, b) }

func BenchmarkIterativeFib1(b *testing.B)  { benchmarkIterativeFibFor(1, b) }
func BenchmarkIterativeFib2(b *testing.B)  { benchmarkIterativeFibFor(2, b) }
func BenchmarkIterativeFib5(b *testing.B)  { benchmarkIterativeFibFor(5, b) }
func BenchmarkIterativeFib10(b *testing.B) { benchmarkIterativeFibFor(10, b) }
func BenchmarkIterativeFib20(b *testing.B) { benchmarkIterativeFibFor(20, b) }
func BenchmarkIterativeFib40(b *testing.B) { benchmarkIterativeFibFor(40, b) }

func benchmarkRecursiveFibFor(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		RecursiveFib(i)
	}
}

func benchmarkIterativeFibFor(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		IterativeFib(i)
	}
}
