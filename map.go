package seq

import (
	"context"
	"iter"
	"sync"
)

type mapfunc[T, U any] func(T) (U, error)

type mapSeq[T, U any] struct {
	f mapfunc[T, U]
	i iter.Seq2[T, error]
}

func Map[T, U any](s Seq[T], f mapfunc[T, U]) Seq[U] {
	return mapSeq[T, U]{f, s.Iterator()}
}

func (s mapSeq[T, U]) Iterator() iter.Seq2[U, error] {
	return func(yield func(U, error) bool) {
		var r U

		for v, err := range s.i {

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

type mapAsyncSeq[T, U any] struct {
	f mapfunc[T, U]
	i iter.Seq2[T, error]
}

func MapAsync[T, U any](s Seq[T], f mapfunc[T, U]) Seq[U] {
	return mapAsyncSeq[T, U]{f, s.Iterator()}
}

func (s mapAsyncSeq[T, U]) Iterator() iter.Seq2[U, error] {
	return func(yield func(U, error) bool) {
		var r U
		var wg sync.WaitGroup
		var l sync.Mutex

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// channel to replace sync.Mutex
		// optional: fail fast mode

		for v, err := range s.i {
			wg.Add(1)

			if err != nil {
				if !yield(r, err) {
					return
				}
			}

			go func(v T) {
				defer wg.Done()

				r, err := s.f(v)

				l.Lock()
				defer l.Unlock()

				select {
				case <-ctx.Done():
					return
				default:
					if !yield(r, err) {
						cancel()
					}
				}
			}(v)
		}

		wg.Wait()
	}
}
