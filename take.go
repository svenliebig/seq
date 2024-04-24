package seq

type takeSeq[T any] struct {
	i Iterator[T]
	n int
}

// Takes the first n elements from a sequence.
func Take[T any](s Seq[T], n int) Seq[T] {
	return takeSeq[T]{s.Iterator(), n}
}

func (s takeSeq[T]) Iterator() Iterator[T] {
	return func(yield func(T, error) bool) {
		for v, err := range s.i {
			if err != nil {
				if !yield(v, err) {
					return
				}
			}

			if s.n == 0 {
				return
			}

			s.n--

			if !yield(v, nil) {
				return
			}
		}
	}
}
