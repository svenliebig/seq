package seq

import (
	"iter"
)

type Iterator[T any] iter.Seq2[T, error]

type Seq[T any] interface {
	Iterator() iter.Seq2[T, error]
}
