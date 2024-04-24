package seq

type mapSeq[T, U any] struct {
	f func(T) U
	i Iterator[T]
}

func Map[T, U any](s Seq[T], f func(T) U) Seq[U] {
	return mapSeq[T, U]{f, s.Iterator()}
}

func (s mapSeq[T, U]) Iterator() Iterator[U] {
	return func(yield func(U, error) bool) {
		for v, err := range s.i {
			if err != nil {
				if !yield(s.f(v), err) {
					return
				}
			}

			if !yield(s.f(v), nil) {
				return
			}
		}
	}
}
