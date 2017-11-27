# cgobench
Simple benchmark of cgo calling overhead

Results from my machine:
------------------------

    scott@saruman:~/Dropbox/go/src/github.com/nu11ptr/cgobench$ go test -bench=.
    goos: linux
    goarch: amd64
    pkg: github.com/nu11ptr/cgobench
    BenchmarkCGo0Args-8             20000000                64.6 ns/op
    BenchmarkCGo1IntArg-8           20000000                62.5 ns/op
    BenchmarkCGo3PtrArgs-8          20000000                65.3 ns/op
    Benchmark0Args-8                2000000000               0.25 ns/op
    Benchmark1IntArg-8              2000000000               0.25 ns/op
    Benchmark3PtrArgs-8             2000000000               0.25 ns/op
    PASS
    ok      github.com/nu11ptr/cgobench     5.654s
