package seq

import (
	"fmt"
	"testing"
	"time"
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

func TestAsyncExecutionTime(t *testing.T) {
	t.Run("should run in a reasonable amount of time", func(t *testing.T) {
		_, err := Collect(
			MapAsync(
				Int(0, 5000),
				func(i int) (int, error) {
					time.Sleep(1 * time.Second)
					return i, nil
				},
			),
		)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	})
}

func TestAsyncError(t *testing.T) {
	t.Run("should handle errors correctly", func(t *testing.T) {
		_, err := Collect(
			MapAsync(
				Int(0, 5000),
				func(i int) (int, error) {
					return 0, fmt.Errorf("error")
				},
			),
		)

		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}

// probably not possible to keep the order?
func TestAsyncOrder(t *testing.T) {
	t.Run("should map async over some numbers", func(t *testing.T) {
		r, err := Collect(
			MapAsync(
				Int(0, 10),
				func(i int) (int, error) {
					// wait for 1 second
					fmt.Println("Processing", i)
					time.Sleep(1 * time.Second)

					return i * 2, nil
				},
			),
		)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		expected := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
		for i, v := range r {
			if v != expected[i] {
				t.Errorf("Expected %d, got %d", expected[i], v)
			}
		}
	})
}
