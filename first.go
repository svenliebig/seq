package seq

import "iter"

type firstSeq[T any] struct {
	p predicate[T]
	i iter.Seq2[T, error]
}

// First returns the first element in the sequence that satisfies the predicate.
func First[T any](s Seq[T], p predicate[T]) Seq[T] {
	return firstSeq[T]{p, s.Iterator()}
}

func (s firstSeq[T]) Iterator() iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for v, err := range s.i {
			if err != nil {
				if !yield(v, err) {
					return
				}
			}

			r, err := s.p(v)

			if err != nil {
				if !yield(v, err) {
					return
				}
			}

			if r {
				yield(v, nil)
				return
			}
		}
	}
}
