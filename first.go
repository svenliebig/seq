package seq

type firstSeq[T any] struct {
	p func(T) bool
	i Iterator[T]
}

// First returns the first element in the sequence that satisfies the predicate.
func First[T any](s Seq[T], p func(T) bool) Seq[T] {
	return firstSeq[T]{p, s.Iterator()}
}

func (s firstSeq[T]) Iterator() Iterator[T] {
	return func(yield func(T, error) bool) {
		for v, err := range s.i {
			if err != nil {
				if !yield(v, err) {
					return
				}
			}

			if s.p(v) {
				yield(v, nil)
				return
			}
		}
	}
}
