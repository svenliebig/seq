package seq

import (
	"fmt"
	"testing"
)

func TestSeq(t *testing.T) {
	t.Run("IntSeq", func(t *testing.T) {
		it := intSeq{start: 0, end: 4}.Iterator()

		result := ""
		for i := range it {
			result += fmt.Sprint(i)
		}

		if result != "01234" {
			t.Errorf("Expected 01234, got %s", result)
		}
	})
}
