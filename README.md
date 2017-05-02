# This project benchmarks the [github.com/adam-hanna/sessions](https://github.com/adam-hanna/sessions) sessions framework against the [github.com/gorilla/sessions](https://github.com/gorilla/sessions) framework.

## [github.com/adam-hanna/sessions](https://github.com/adam-hanna/sessions)
As stated in the package, my benchmark results are as follows:

~~~ bash
$ (cd benchmark && go test -bench=.)

setting up benchmark tests
BenchmarkBaseServer-2              20000             72479 ns/op
BenchmarkValidSession-2            10000            151650 ns/op
PASS
shutting down benchmark tests
ok      github.com/adam-hanna/sessions/benchmark        3.727s
~~~

## [github.com/gorilla/sessions](https://github.com/gorilla/sessions) 
The gorilla sessions server was setup as shown in main.go. The server was started before performing the benchmarks. The benchmarks were run on the same machine, on the same day as the benchmarks given above (FWTW).

~~~ bash
$ (cd benchmark && go test -bench=.)

BenchmarkValidSession-2             5000            310136 ns/op
PASS
ok      github.com/adam-hanna/sessions-comparison/gorilla-sessions/benchmark    1.593s
~~~