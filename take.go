package seq

import (
	"iter"
)

type takeSeq[T any] struct {
	i iter.Seq2[T, error]
	n int
}

// Takes the first n elements from a sequence.
func Take[T any](s Seq[T], n int) Seq[T] {
	return takeSeq[T]{s.Iterator(), n}
}

func (s takeSeq[T]) Iterator() iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		next, stop := iter.Pull2(s.i)
		defer stop()

		for s.n > 0 {
			v, err, valid := next()
			s.n--

			if !valid {
				return
			}

			if err != nil {
				if !yield(v, err) {
					return
				}
			} else if !yield(v, nil) {
				return
			}
		}
	}
}
