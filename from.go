package seq

import "iter"

type fromSeq[T any] struct {
	s []T
}

func From[T any](s []T) Seq[T] {
	return fromSeq[T]{s}
}

func (s fromSeq[T]) Iterator() iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for _, v := range s.s {
			if !yield(v, nil) {
				return
			}
		}
	}
}
