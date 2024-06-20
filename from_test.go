package seq

import (
	"fmt"
	"testing"
)

func TestFrom(t *testing.T) {
	t.Run("should create an int sequence", func(t *testing.T) {
		it := From([]int{1, 2, 3}).Iterator()

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

	t.Run("should create a string sequence", func(t *testing.T) {
		it := From([]string{"a", "b", "c"}).Iterator()

		result := ""
		for i, err := range it {
			if err != nil {
				t.Errorf("didn't expect any error, but got: %v", err)
				return
			}

			result += i
		}

		if result != "abc" {
			t.Errorf("Expected abc, got %s", result)
		}
	})
}
