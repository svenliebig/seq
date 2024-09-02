package seq

import (
	"context"
	"iter"
	"sync"
)

type mapfunc[T, U any] func(T) (U, error)

type mapSeq[I, O any] struct {
	f mapfunc[I, O]
	i iter.Seq2[I, error]
}

func Map[I, O any](s Seq[I], f mapfunc[I, O]) Seq[O] {
	return mapSeq[I, O]{f, s.Iterator()}
}

func (s mapSeq[I, O]) Iterator() iter.Seq2[O, error] {
	return func(yield func(O, error) bool) {
		var r O

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

// BenchmarkMapAsync5000Entries-8               944           1289784 ns/op          650616 B/op        10053 allocs/op
func (s mapAsyncSeq[T, U]) IteratorOld() iter.Seq2[U, error] {
	return func(yield func(U, error) bool) {
		var r U
		var wg sync.WaitGroup
		var l sync.Mutex

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

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

// BenchmarkMapAsync5000Entries-8               112          10528610 ns/op          727895 B/op        11679 allocs/op
func (s mapAsyncSeq[T, U]) Iterator() iter.Seq2[U, error] {
	return func(yield func(U, error) bool) {
		var r U
		var wg sync.WaitGroup

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		c := make(chan struct {
			v U
			e error
		})

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

				select {
				case <-ctx.Done():
					return
				case c <- struct {
					v U
					e error
				}{r, err}: // send to channel
				}
			}(v)
		}

		go func() {
			wg.Wait()
			close(c)
		}()

		for v := range c {
			if !yield(v.v, v.e) {
				cancel()
				break
			}
		}

	}
}
