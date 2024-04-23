package seq

import (
	"iter"
)

type mapSeq[T, U any] struct {
	f func(T) U
	i iter.Seq[T]
}

func Map[T, U any](s Seq[T], f func(T) U) Seq[U] {
	return mapSeq[T, U]{f, s.Iterator()}
}

func (s mapSeq[T, U]) Iterator() iter.Seq[U] {
	return func(yield func(U) bool) {
		for e := range s.i {
			if !yield(s.f(e)) {
				return
			}
		}
	}
}
