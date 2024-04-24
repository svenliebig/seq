package seq

import "iter"

type emptySeq struct{}

// Creates an empty sequence.
func Empty() Seq[any] {
	return emptySeq{}
}

func (s emptySeq) Iterator() iter.Seq2[any, error] {
	return func(yield func(any, error) bool) {}
}
