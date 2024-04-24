package seq

import (
	"fmt"
	"testing"
)

func TestTake(t *testing.T) {
	t.Run("should take the first three integers", func(t *testing.T) {
		it := Take(
			Int(0, 5),
			3,
		).Iterator()

		var result string
		for v, err := range it {
			if err != nil {
				t.Errorf("Expected nil, got %v", err)
			}

			result += fmt.Sprint(v)
		}

		expected := "012"

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
