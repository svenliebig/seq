# notes & insights

This has 2 allocations:

```go
func BenchmarkFilterTest(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = Filter(
            Int(0, 1000),
            func(i int) bool {
                return i%2 == 0
            },
        )
    }
    b.ReportAllocs()
}
```

And this created 4 allocations:

```go
func BenchmarkFilterTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Filter(
			Int(0, 1000),
			func(i int) bool {
				return i%2 == 0
			},
		).Iterator()
	}
	b.ReportAllocs()
}
```
