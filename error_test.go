package seq

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	t.Run("should return an error as first alue", func(t *testing.T) {
		s := Error[any](errors.New("error"))

		for _, err := range s.Iterator() {
			if err == nil {
				t.Errorf("expected an error, got %v", nil)
			}

			if err.Error() != "error" {
				t.Errorf("expected an error, got %v", err)
			}
		}
	})
}
