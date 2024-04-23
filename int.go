package seq

import (
	"iter"
)

var IntSeq Seq[int] = intSeq{}

type intSeq struct {
	start, end int
	reverse    bool
}

// Int returns a sequence of integers from start to end.
//
// The sequence is inclusive of start and end.
//
// If start is greater than end, the sequence will decrement from start to end, otherwise it will increment.
func Int(start, end int) Seq[int] {
	r := start > end
	return intSeq{start, end, r}
}

func (s intSeq) Iterator() iter.Seq[int] {
	return func(yield func(int) bool) {
		if s.reverse {
			for i := s.start; i >= s.end; i-- {
				if !yield(i) {
					return
				}
			}
		}

		for i := s.start; i <= s.end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}
