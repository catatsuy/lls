I use Docker on M1 Mac.

```
goos: linux
goarch: arm64
pkg: github.com/catatsuy/lls/bench
BenchmarkUseStringCast-8        17582245                67.68 ns/op          256 B/op          1 allocs/op
BenchmarkUseUnsafeString-8      1000000000               1.000 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/catatsuy/lls/bench   2.193s
```
