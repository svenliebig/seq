package seq

import "iter"

type skipSeq[T any] struct {
	n int
	i iter.Seq2[T, error]
}

// Skips the first n elements from a sequence.
func Skip[T any](s Seq[T], n int) Seq[T] {
	return skipSeq[T]{n, s.Iterator()}
}

func (s skipSeq[T]) Iterator() iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for v, err := range s.i {
			if err != nil {
				if !yield(v, err) {
					return
				}
			}

			if s.n > 0 {
				s.n--
				continue
			}

			if !yield(v, nil) {
				return
			}
		}
	}
}
