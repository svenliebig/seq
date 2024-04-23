package seq

func Reduce[T, U any](s Seq[T], f func(U, T) U) U {
	var acc U

	for e := range s.Iterator() {
		acc = f(acc, e)
	}

	return acc
}
