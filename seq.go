package seq

import (
	"iter"
)

type Seq[T any] interface {
	Iterator() iter.Seq[T]
}
