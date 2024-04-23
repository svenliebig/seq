package seq

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	s := Filter(
		Int(0, 10),
		func(i int) bool {
			return i%2 == 0
		},
	)

	res := ""
	expected := "0246810"

	for v := range s.Iterator() {
		res += fmt.Sprint(v)
	}

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}
