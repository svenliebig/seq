package seq

import "iter"

type filterSeq[T any] struct {
	f func(T) bool
	s Seq[T]
}

func Filter[T any](s Seq[T], f func(T) bool) Seq[T] {
	return filterSeq[T]{f, s}
}

func (s filterSeq[T]) Iterator() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range s.s.Iterator() {
			if !s.f(e) {
				continue
			}

			if !yield(e) {
				return
			}
		}
	}
}
