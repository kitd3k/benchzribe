package benchmarks

import (
	"testing"
)

func BenchmarkSimpleOperation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for j := 0; j < 1000; j++ {
			sum += j
		}
	}
}

func BenchmarkStringConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < 100; j++ {
			s += "test"
		}
	}
}

func BenchmarkSliceOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 0, 1000)
		for j := 0; j < 1000; j++ {
			slice = append(slice, j)
		}
	}
} 