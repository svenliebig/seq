package seq

import (
	"fmt"
	"testing"
)

func TestSkip(t *testing.T) {
	t.Run("should skip the first 4 integers and result in 5", func(t *testing.T) {
		it := Skip(
			Int(1, 5),
			4,
		).Iterator()

		var result int
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

	t.Run("should skip the first 2 and iterate over the other 3", func(t *testing.T) {
		it := Skip(
			Int(1, 5),
			2,
		).Iterator()

		result := ""
		for v := range it {
			result += fmt.Sprint(v)
		}

		expected := "345"

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
