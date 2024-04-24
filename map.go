package seq

type mapfunc[T, U any] func(T) (U, error)

type mapSeq[T, U any] struct {
	f mapfunc[T, U]
	i Iterator[T]
}

func Map[T, U any](s Seq[T], f mapfunc[T, U]) Seq[U] {
	return mapSeq[T, U]{f, s.Iterator()}
}

func (s mapSeq[T, U]) Iterator() Iterator[U] {
	return func(yield func(U, error) bool) {
		for v, err := range s.i {
			var r U

			if err != nil {
				if !yield(r, err) {
					return
				}
			}

			if !yield(s.f(v)) {
				return
			}
		}
	}
}
