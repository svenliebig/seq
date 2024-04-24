package seq

import (
	"fmt"
	"testing"
)

func TestBatch(t *testing.T) {
	t.Run("should batch into an int pairs", func(t *testing.T) {
		s := Batch(
			Int(0, 9),
			2,
		)

		res := ""
		expected := "01,23,45,67,89,"

		for v, err := range s.Iterator() {
			if err != nil {
				t.Errorf("Expected no error, got %s", err)
			}

			res += fmt.Sprintf("%d%d,", v[0], v[1])
		}

		if res != expected {
			t.Errorf("Expected %s, got %s", expected, res)
		}
	})
}
