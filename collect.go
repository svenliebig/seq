package seq

func Collect[T any](s Seq[T]) []T {
	var res []T

	for e := range s.Iterator() {
		res = append(res, e)
	}

	return res
}
