package seq

import (
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("should map integer values to strings", func(t *testing.T) {
		s := Map(
			Int(3, 0),
			func(i int) string {
				switch i {
				case 0:
					return "Zero"
				case 1:
					return "One"
				case 2:
					return "Two"
				case 3:
					return "Three"
				default:
					return "Unknown"
				}
			},
		)

		res := ""
		expected := "ThreeTwoOneZero"

		for v := range s.Iterator() {
			res += v
		}

		if res != expected {
			t.Errorf("Expected %s, got %s", expected, res)
		}
	})
}
