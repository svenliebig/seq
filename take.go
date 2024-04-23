package seq

import (
	"iter"
)

type takeSeq[T any] struct {
	i iter.Seq[T]
	n int
}

// Takes the first n elements from a sequence.
func Take[T any](s Seq[T], n int) Seq[T] {
	return takeSeq[T]{s.Iterator(), n}
}

func (s takeSeq[T]) Iterator() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range s.i {
			if s.n == 0 {
				return
			}

			s.n--

			if !yield(e) {
				return
			}
		}
	}
}
