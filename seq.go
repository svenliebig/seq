package seq

import (
	"errors"
	"iter"
)

var (
	ErrEndOfSeq = errors.New("end of sequence")
)

type Seq[T any] interface {
	Iterator() iter.Seq[T]
}

