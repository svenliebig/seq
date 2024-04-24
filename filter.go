package seq

type predicate[T any] func(T) (bool, error)

type filterSeq[T any] struct {
	p predicate[T]
	i Iterator[T]
}

func Filter[T any](s Seq[T], p predicate[T]) Seq[T] {
	return filterSeq[T]{p, s.Iterator()}
}

func (s filterSeq[T]) Iterator() Iterator[T] {
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

			if !r {
				continue
			}

			if !yield(v, nil) {
				return
			}
		}
	}
}
