package seq

func Collect[T any](s Seq[T]) ([]T, error) {
	var res []T

	for v, err := range s.Iterator() {
		if err != nil {
			return nil, err
		}

		res = append(res, v)
	}

	return res, nil
}
