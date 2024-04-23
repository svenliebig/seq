package seq

import (
	"fmt"
	"testing"
)

func TestIntSeq(t *testing.T) {
	t.Run("should create a sequence from 0 to 4", func(t *testing.T) {
		it := Int(0, 4).Iterator()

		result := ""
		for i := range it {
			result += fmt.Sprint(i)
		}

		if result != "01234" {
			t.Errorf("Expected 01234, got %s", result)
		}
	})

	t.Run("should create a sequence from 4 to 0", func(t *testing.T) {
		it := Int(4, 0).Iterator()

		result := ""
		for i := range it {
			result += fmt.Sprint(i)
		}

		if result != "43210" {
			t.Errorf("Expected 43210, got %s", result)
		}
	})
}

func BenchmarkIntSeq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it := Int(0, 1000).Iterator()

		target := make([]int, 0, 1000)

		for i := range it {
			target = append(target, i)
		}
	}
}

func BenchmarkIntNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		target := make([]int, 0, 1000)

		for i := 0; i <= 1000; i++ {
			target = append(target, i)
		}
	}
}
