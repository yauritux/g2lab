# How Benchmarks Work 
Each benchmark is run b.N times until it takes longer than 1 second.

b.N starts at 1, if the benchmark completes is under 1 second, b.N is increased and the benchmark run again.

b.N increases in the approximate sequence; 1, 2, 3, 5, 10, 20, 30, 50, 100, ...

```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/yauritux/benchmarking
BenchmarkFib20-8   	   30000	     45151 ns/op
PASS
ok  	github.com/yauritux/benchmarking	1.826s
```

*Beware*: below the microsecond mark you will start to see the relativistic effects of instruction reordering and code alignment.

Run benchmarks longer to get more accuracy; `go test -benchtime=10s`

*Tip*: If this is required, codify it in a `Makefile` so everyone is comparing apples to apples.

# Comparing Benchmarks

For repeatable results, you should run benchmark multiple times.

You can do this manually, or use the -count flag.

```
$ go test -bench=. -count=10 | tee old.txt
goos: darwin
goarch: amd64
pkg: github.com/yauritux/benchmarking
BenchmarkFib20-8           30000             44822 ns/op
BenchmarkFib20-8           30000             44966 ns/op
BenchmarkFib20-8           30000             44948 ns/op
BenchmarkFib20-8           30000             45058 ns/op
BenchmarkFib20-8           30000             44980 ns/op
BenchmarkFib20-8           30000             44827 ns/op
BenchmarkFib20-8           30000             44795 ns/op
BenchmarkFib20-8           30000             44983 ns/op
BenchmarkFib20-8           30000             44956 ns/op
BenchmarkFib20-8           30000             44733 ns/op
PASS
ok      github.com/yauritux/benchmarking        18.031s
```
After refactor the Fib function (in fibnew.go), run the benchmark again:

```
$ go test -bench=. -count=10 | tee new.txt
goos: darwin
goarch: amd64
pkg: github.com/yauritux/benchmarking
BenchmarkFib20-8   	   50000	     28240 ns/op
BenchmarkFib20-8   	   50000	     28236 ns/op
BenchmarkFib20-8   	   50000	     28782 ns/op
BenchmarkFib20-8   	   50000	     28478 ns/op
BenchmarkFib20-8   	   50000	     28091 ns/op
BenchmarkFib20-8   	   50000	     28350 ns/op
BenchmarkFib20-8   	   50000	     28252 ns/op
BenchmarkFib20-8   	   50000	     28710 ns/op
BenchmarkFib20-8   	   50000	     28322 ns/op
BenchmarkFib20-8   	   50000	     28261 ns/op
PASS
ok  	github.com/yauritux/benchmarking	17.082s
```

Determining the performance delta between two set of benchmarks can be tedious and error prone.

Tools like rsc.io/benchstat are useful for comparing results

```
$ go get -u rsc.io/benchstat
$ benchstat {old,new}.txt
name     old time/op  new time/op  delta
Fib20-8  44.9µs ± 0%  28.4µs ± 1%  -36.82%  (p=0.000 n=10+10)
```

*Tip*: p values above 0.05 are suspect, increase -count to add more samples.


