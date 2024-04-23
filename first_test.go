package seq

import (
	"testing"
)

func TestFirst(t *testing.T) {
	t.Run("should get the 5 in the integer sequence", func(t *testing.T) {
		it := First(Int(1, 10), func(i int) bool {
			return i%5 == 0
		}).Iterator()

		result := 0
		iterations := 0

		for v := range it {
			result = v
			iterations++
		}

		if result != 5 {
			t.Errorf("Expected 5, got %d", result)
		}

		if iterations != 1 {
			t.Errorf("Expected 1 iteration, got %d", iterations)
		}
	})
}
