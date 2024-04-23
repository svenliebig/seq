package seq

import (
	"fmt"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("should sum integers", func(t *testing.T) {
		it := Reduce(
			Int(0, 5),
			func(acc, v int) int {
				return acc + v
			},
		).Iterator()

		var result int
		iterations := 0
		for v := range it {
			result = v
			iterations++
		}

		if result != 15 {
			t.Errorf("Expected 15, got %d", result)
		}

		if iterations != 1 {
			t.Errorf("Expected 1 iteration, got %d", iterations)
		}
	})

	t.Run("should append integers to a string", func(t *testing.T) {
		it := Reduce(
			Int(0, 5),
			func(acc string, v int) string {
				return acc + fmt.Sprint(v)
			},
		).Iterator()

		var result string
		iterations := 0
		for v := range it {
			result = v
			iterations++
		}

		if result != "012345" {
			t.Errorf("Expected 012345, got %s", result)
		}

		if iterations != 1 {
			t.Errorf("Expected 1 iteration, got %d", iterations)
		}
	})

}

func BenchmarkReduce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for v := range Reduce(
			Int(0, 1000),
			func(acc, v int) int {
				return acc + v
			},
		).Iterator() {
			_ = v
		}
	}
	b.ReportAllocs()
}

func BenchmarkReduceNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var result int
		for v := 0; v <= 1000; v++ {
			result += v
		}
	}
}
