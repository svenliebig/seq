package seq

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("should filter even numbers", func(t *testing.T) {
		s := Filter(
			Int(0, 10),
			func(i int) bool {
				return i%2 == 0
			},
		)

		res := ""
		expected := "0246810"

		for v := range s.Iterator() {
			res += fmt.Sprint(v)
		}

		if res != expected {
			t.Errorf("Expected %s, got %s", expected, res)
		}
	})
}

func BenchmarkFilterEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Filter(
			Empty(),
			func(i any) bool {
				return true
			},
		)
	}
	b.ReportAllocs()
}

func BenchmarkFilterEmptyIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Filter(
			Empty(),
			func(i any) bool {
				return true
			},
		).Iterator()
	}
	b.ReportAllocs()
}

func BenchmarkFilterTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Filter(
			Int(0, 1000),
			func(i int) bool {
				return i%2 == 0
			},
		).Iterator()
	}
	b.ReportAllocs()
}

func BenchmarkFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := Filter(
			Int(0, 1000),
			func(i int) bool {
				return i%2 == 0
			},
		)

		target := make([]int, 0, 500)

		for v := range s.Iterator() {
			target = append(target, v)
		}
	}
	b.ReportAllocs()
}

func BenchmarkNativeFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		target := make([]int, 0, 500)

		for i := 0; i <= 1000; i++ {
			if i%2 == 0 {
				target = append(target, i)
			}
		}
	}
	b.ReportAllocs()
}
