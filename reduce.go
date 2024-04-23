package seq

import (
	"iter"
)

type reduceSeq[T, U any] struct {
	f func(U, T) U
	s iter.Seq[T]
}

func Reduce[T, U any](s Seq[T], f func(U, T) U) Seq[U] {
	return &reduceSeq[T, U]{f, s.Iterator()}
}

func (s reduceSeq[T, U]) Iterator() iter.Seq[U] {
	return func(yield func(U) bool) {
		var acc U

		for e := range s.s {
			acc = s.f(acc, e)
		}

		if !yield(acc) {
			return
		}
	}
}
