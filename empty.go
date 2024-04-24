package seq

type emptySeq struct{}

// Creates an empty sequence.
func Empty() Seq[any] {
	return emptySeq{}
}

func (s emptySeq) Iterator() Iterator[any] {
	return func(yield func(any, error) bool) {}
}
