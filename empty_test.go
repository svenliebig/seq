package seq

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	t.Run("should not call any iteration", func(t *testing.T) {
		s := Empty()

		for v := range s.Iterator() {
			t.Errorf("Expected no iteration, got %v", v)
		}
	})
}

func BenchmarkEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Empty()
	}
	b.ReportAllocs()
}

func BenchmarkEmptyIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Empty().Iterator()
	}
	b.ReportAllocs()
}
