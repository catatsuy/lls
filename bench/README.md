I use Docker on M1 Mac.

```
goos: linux
goarch: arm64
pkg: github.com/catatsuy/lls/bench
BenchmarkUseStringCast-4         3001924               369.4 ns/op           512 B/op          2 allocs/op
BenchmarkUseUnsafeString-4       3652689               329.4 ns/op           256 B/op          1 allocs/op
PASS
ok      github.com/catatsuy/lls/bench   3.063s

goos: linux
goarch: arm64
pkg: github.com/catatsuy/lls/bench
BenchmarkUseStringCast-4         3254386               365.3 ns/op           512 B/op          2 allocs/op
BenchmarkUseUnsafeString-4       3646972               328.3 ns/op           256 B/op          1 allocs/op
PASS
ok      github.com/catatsuy/lls/bench   3.102s

goos: linux
goarch: arm64
pkg: github.com/catatsuy/lls/bench
BenchmarkUseStringCast-4         3254906               368.6 ns/op           512 B/op          2 allocs/op
BenchmarkUseUnsafeString-4       3656787               328.3 ns/op           256 B/op          1 allocs/op
PASS
ok      github.com/catatsuy/lls/bench   3.116s
```
