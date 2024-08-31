package seq

import "iter"

type errorSeq[T any] struct {
	err error
}

// creates an error sequence, this sequence will yield the given error and no value.
func Error[T any](err error) Seq[T] {
	return errorSeq[T]{
		err: err,
	}
}

func (s errorSeq[T]) Iterator() iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		var r T
		yield(r, s.err)
	}
}
