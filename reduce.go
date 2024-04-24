package seq

func Reduce[T, U any](s Seq[T], f func(U, T) (U, error)) (U, error) {
	var acc U

	for v, err := range s.Iterator() {
		if err != nil {
			return acc, err
		}

		acc, err = f(acc, v)

		if err != nil {
			return acc, err
		}
	}

	return acc, nil
}
