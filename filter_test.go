package seq

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("should filter even numbers", func(t *testing.T) {
		s := Filter(
			Int(0, 10),
			func(i int) (bool, error) {
				return i%2 == 0, nil
			},
		)

		res := ""
		expected := "0246810"

		for v, err := range s.Iterator() {
			if err != nil {
				t.Error(err)
				return
			}

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
			func(i any) (bool, error) {
				return true, nil
			},
		)
	}
	b.ReportAllocs()
}

// PUZZLE: this allocates 2 times, while the previous benchmark allocates 0 times
func BenchmarkFilterEmptyIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Filter(
			Empty(),
			func(i any) (bool, error) {
				return true, nil
			},
		).Iterator()
	}
	b.ReportAllocs()
}

// 2 allocs
func BenchmarkFilterIntDeclaration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Filter(
			Int(0, 1000),
			func(i int) (bool, error) {
				return i%2 == 0, nil
			},
		)
	}
	b.ReportAllocs()
}

// 4 allocs
func BenchmarkFilterIntDeclarationIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Filter(
			Int(0, 1000),
			func(i int) (bool, error) {
				return i%2 == 0, nil
			},
		).Iterator()
	}
	b.ReportAllocs()
}

func dummy(i int) {
}

// 9 allocs
func BenchmarkFilterIntDeclarationIteratorRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it := Filter(
			Int(0, 1000),
			func(i int) (bool, error) {
				return i%2 == 0, nil
			},
		).Iterator()

		for v, _ := range it {
			dummy(v)
		}
	}
	b.ReportAllocs()
}

// PUZZLE: 11 allocs/op
func BenchmarkFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 4 allocs
		it := Filter(
			Int(0, 1000),
			func(i int) (bool, error) {
				return i%2 == 0, nil
			},
		).Iterator()

		// 1 alloc
		target := make([]int, 0, 501)

		// 5 allocs
		for v, _ := range it {
			target = append(target, v)
		}

		// ?? last alloc
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
