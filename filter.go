package seq

import "iter"

type filterSeq[T any] struct {
	f func(T) bool
	i iter.Seq[T]
}

func Filter[T any](s Seq[T], f func(T) bool) Seq[T] {
	return filterSeq[T]{f, s.Iterator()}
}

func (s filterSeq[T]) Iterator() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range s.i {
			if !s.f(e) {
				continue
			}

			if !yield(e) {
				return
			}
		}
	}
}
