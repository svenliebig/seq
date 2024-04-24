package seq

import (
	"iter"
)

type batchSeq[T any] struct {
	s iter.Seq2[T, error]
	n int
}

// Batch returns a new sequence that puts n elements of the input sequence into a slice.
func Batch[T any](s Seq[T], n int) Seq[[]T] {
	return &batchSeq[T]{s.Iterator(), n}
}

func (s *batchSeq[T]) Iterator() iter.Seq2[[]T, error] {
	return func(yield func([]T, error) bool) {
		n, stop := iter.Pull2(s.s)
		defer stop()

		for e, err, valid := n(); valid; {
			batch := make([]T, s.n)

			for i := 0; i < s.n; i++ {
				batch[i] = e
				e, err, valid = n()
			}

			if !yield(batch, err) {
				return
			}
		}
	}
}
