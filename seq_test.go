package seq

import (
	"fmt"
	"testing"
)

func TestSeq(t *testing.T) {
	t.Run("should correctly raise errors in a Collect, Filter chain", func(t *testing.T) {
		_, err := Collect(
			Filter(
				Int(0, 3),
				func(i int) (bool, error) {
					if i == 2 {
						return false, fmt.Errorf("error")
					}

					return true, nil
				},
			))

		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})

	t.Run("should correctly raise errors in a Collect, First, Skip, Take, Filter chain", func(t *testing.T) {
		_, err := Collect(
			First(
				Skip(
					Take(
						Filter(
							Int(0, 10),
							func(i int) (bool, error) {
								if i == 8 {
									return false, fmt.Errorf("error")
								}

								return true, nil
							},
						),
						8,
					),
					5,
				),
				func(i int) (bool, error) {
					return i == 8, nil
				},
			),
		)

		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})

	t.Run("should correctly raise errors in a Reduce, Map chain", func(t *testing.T) {
		_, err := Reduce(
			Map(
				Int(0, 10),
				func(i int) (string, error) {
					if i == 5 {
						return "", fmt.Errorf("error")
					}

					return fmt.Sprintf("%d", i), nil
				},
			),
			func(acc string, i string) (string, error) {
				return acc + i, nil
			},
		)

		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}

func BenchmarkSeqCreateMapFilterCollect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Collect(
			Map(
				Filter(
					Int(0, 10000),
					func(i int) (bool, error) { return i%2 == 0, nil },
				),
				func(i int) (int, error) { return i * 2, nil },
			),
		)
	}
}

func BenchmarkSeqCreatemapFilterCollectNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var result []int
		for i := 0; i < 10000; i++ {
			if i%2 == 0 {
				result = append(result, i*2)
			}
		}
	}
}
