package seq

import (
	"iter"
)

type firstSeq[T any] struct {
	p func(T) bool
	i iter.Seq[T]
}

// First returns the first element in the sequence that satisfies the predicate.
func First[T any](s Seq[T], p func(T) bool) Seq[T] {
	return firstSeq[T]{p, s.Iterator()}
}

func (s firstSeq[T]) Iterator() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range s.i {
			if s.p(e) {
				yield(e)
				return
			}
		}
	}
}
