package seq

import (
	"iter"
)

type Seq[T any] interface {
	Iterator() iter.Seq2[T, error]
}
