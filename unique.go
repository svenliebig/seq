package seq

import "iter"

type uniqueSeq[T comparable] struct {
	i iter.Seq2[T, error]
}

func Unique[T comparable](s Seq[T]) Seq[T] {
	return uniqueSeq[T]{s.Iterator()}
}

func (s uniqueSeq[T]) Iterator() iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		a := make([]T, 0)

		for v, err := range s.i {
			if err != nil {
				yield(v, err)
				return
			}

			if contains(a, v) {
				continue
			}

			a = append(a, v)

			if !yield(v, nil) {
				return
			}
		}
	}
}

func contains[T comparable](s []T, c T) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}

	return false
}
