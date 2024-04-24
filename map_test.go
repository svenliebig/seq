package seq

import (
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("should map integer values to strings", func(t *testing.T) {
		s := Map(
			Int(3, 0),
			func(i int) (string, error) {
				switch i {
				case 0:
					return "Zero", nil
				case 1:
					return "One", nil
				case 2:
					return "Two", nil
				case 3:
					return "Three", nil
				default:
					return "Unknown", nil
				}
			},
		)

		res := ""
		expected := "ThreeTwoOneZero"

		for v, _ := range s.Iterator() {
			res += v
		}

		if res != expected {
			t.Errorf("Expected %s, got %s", expected, res)
		}
	})
}
