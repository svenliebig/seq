package seq

import (
	"testing"
)

func BenchmarkSeqCreateMapFilterCollect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Collect(
			Map(
				Filter(
					Int(0, 10000),
					func(i int) (bool, error) { return i%2 == 0, nil },
				),
				func(i int) int { return i * 2 },
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
