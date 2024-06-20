package seq

import (
	"fmt"
	"testing"
)

func TestUnique(t *testing.T) {
	t.Run("should not error on empty sequence", func(t *testing.T) {
		s := Unique(
			Empty(),
		)

		for v, err := range s.Iterator() {
			if err != nil {
				t.Errorf("didn't expect any error, but got: %v", err)
				return
			}

			t.Errorf("didn't expect any value, but got: %v", v)
		}
	})

	t.Run("should remove duplicates", func(t *testing.T) {
		it := Unique(From([]int{1, 2, 3, 1, 2, 3, 1, 2, 3})).Iterator()

		result := ""
		for i, err := range it {
			if err != nil {
				t.Errorf("didn't expect any error, but got: %v", err)
				return
			}

			result += fmt.Sprint(i)
		}

		if result != "123" {
			t.Errorf("Expected 123, got %s", result)
		}
	})
}
