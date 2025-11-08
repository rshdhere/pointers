## Benchmark Tests

```bash
go test -bench=. -count 1 -gcflags=-N
```
```bash
goos: linux
goarch: amd64
pkg: github.com/rshdhere/pointers
cpu: AMD Ryzen 5 5500U with Radeon Graphics
BenchmarkPBP-12         330151765                3.623 ns/op
BenchmarkPBV-12           383847              3111 ns/op
PASS
ok      github.com/rshdhere/pointers    2.803s
```

```bash
go tool compile -N -S -l ptr.go
```
