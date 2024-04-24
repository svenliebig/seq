package seq

import (
	"testing"
)

func TestCollect(t *testing.T) {
	t.Run("should collect into a integer slice", func(t *testing.T) {
		r, err := Collect(Int(0, 3))

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		expected := []int{0, 1, 2, 3}

		if len(r) != len(expected) {
			t.Errorf("Expected %v, got %v", expected, r)
		}

		for i, v := range r {
			if v != expected[i] {
				t.Errorf("Expected %v, got %v", expected, r)
			}
		}
	})
}
