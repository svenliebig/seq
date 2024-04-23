package seq

import (
	"iter"
)

type skipSeq[T any] struct {
	n int
	i iter.Seq[T]
}

// Skips the first n elements from a sequence.
func Skip[T any](s Seq[T], n int) Seq[T] {
	return skipSeq[T]{n, s.Iterator()}
}

func (s skipSeq[T]) Iterator() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range s.i {
			if s.n > 0 {
				s.n--
				continue
			}

			if !yield(e) {
				return
			}
		}
	}
}
