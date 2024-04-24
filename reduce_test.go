package seq

import (
	"fmt"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("should sum integers", func(t *testing.T) {
		result, err := Reduce(
			Int(0, 5),
			func(acc, v int) (int, error) {
				return acc + v, nil
			},
		)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if result != 15 {
			t.Errorf("Expected 15, got %d", result)
		}
	})

	t.Run("should append integers to a string", func(t *testing.T) {
		result, err := Reduce(
			Int(0, 5),
			func(acc string, v int) (string, error) {
				return acc + fmt.Sprint(v), nil
			},
		)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if result != "012345" {
			t.Errorf("Expected 012345, got %s", result)
		}
	})
}

func BenchmarkReduce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reduce(
			Int(0, 1000),
			func(acc, v int) (int, error) {
				return acc + v, nil
			},
		)
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
