package seq

import "iter"

type filterMapPredicate[T, U any] func(T) (bool, U, error)

type filterMapSeq[T, U any] struct {
	p filterMapPredicate[T, U]
	i iter.Seq2[T, error]
}

func FilterMap[T, U any](s Seq[T], p filterMapPredicate[T, U]) Seq[U] {
	return filterMapSeq[T, U]{p, s.Iterator()}
}

func (s filterMapSeq[T, U]) Iterator() iter.Seq2[U, error] {
	return func(yield func(U, error) bool) {
		var r U

		for v, err := range s.i {
			if err != nil {
				if !yield(r, err) {
					return
				}
			}

			f, r, err := s.p(v)

			if err != nil {
				if !yield(r, err) {
					return
				}
			}

			if !f {
				continue
			}

			if !yield(r, nil) {
				return
			}
		}
	}
}
