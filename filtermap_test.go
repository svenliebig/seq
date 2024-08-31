package seq

import (
	"testing"
)

func TestFilterMap(t *testing.T) {
	t.Run("should filter one element and map the other", func(t *testing.T) {
		s := FilterMap(
			Int(1, 2),
			func(i int) (bool, string, error) {
				if i == 1 {
					return false, "", nil
				}

				return true, "two", nil
			},
		)

		for v, _ := range s.Iterator() {
			if v != "two" {
				t.Errorf("expected two, got %v", v)
			}
		}
	})
}
