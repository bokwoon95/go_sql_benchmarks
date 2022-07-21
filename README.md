To run these benchmarks, you need a `$POSTGRES_URL` and `$MYSQL_URL`.

The results keep changing, you should run it several times yourself. Even the output below is not necesarily representative of what you will see.

```text
$ go test -bench . -postgres $env:POSTGRES_URL -mysql $env:MYSQL_URL -benchmem -cpuprofile profile
goos: windows
goarch: amd64
pkg: go_sql_benchmarks
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkSqSQLite-16                         429           2761449 ns/op         1223226 B/op      34837 allocs/op
BenchmarkSqCompiledSQLite-16                 442           2711890 ns/op         1222180 B/op      34831 allocs/op
BenchmarkSqlxSQLite-16                       444           2690702 ns/op          662782 B/op      32697 allocs/op
BenchmarkSqPostgres-16                        57          23572089 ns/op          943606 B/op      17431 allocs/op
BenchmarkSqCompiledPostgres-16               100          23488269 ns/op          941744 B/op      17422 allocs/op
BenchmarkSqlxPostgres-16                     100          17330901 ns/op          383360 B/op      15289 allocs/op
BenchmarkSqMySQL-16                          129           8827190 ns/op          987421 B/op      17677 allocs/op
BenchmarkSqCompiledMySQL-16                  698           4783205 ns/op          986431 B/op      17671 allocs/op
BenchmarkSqlxMySQL-16                        794           3018757 ns/op          418782 B/op      13636 allocs/op
PASS
ok      go_sql_benchmarks       18.394s
```
