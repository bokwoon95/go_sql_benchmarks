To run these benchmarks, you need a `$POSTGRES_URL` and `$MYSQL_URL`.

These are the back-to-back runs of the same benchmarks. The results can vary wildly between runs, so even the output below is not necessarily what you will see.


```text
PS C:\Users\cbw\Documents\go_sql_benchmarks> go test -bench . -postgres $env:POSTGRES_URL -mysql $env:MYSQL_URL -benchmem -cpuprofile profile
goos: windows
goarch: amd64
pkg: go_sql_benchmarks
cpu: AMD Ryzen 7 5800H with Radeon Graphics

Benchmark_sq_sqlite-16               435           2653646 ns/op         1223214 B/op      34837 allocs/op
Benchmark_sqlx_sqlite-16             462           2557698 ns/op          662743 B/op      32697 allocs/op

Benchmark_sq_postgres-16             100          17296691 ns/op          942763 B/op      17428 allocs/op
Benchmark_sqlx_postgres-16           100          23485573 ns/op          383384 B/op      15289 allocs/op

Benchmark_sq_mysql-16                255           5264942 ns/op          987399 B/op      17677 allocs/op
Benchmark_sqlx_mysql-16              796           4587407 ns/op          418755 B/op      13636 allocs/op

PASS
ok      go_sql_benchmarks       12.971s

PS C:\Users\cbw\Documents\go_sql_benchmarks> go test -bench . -postgres $env:POSTGRES_URL -mysql $env:MYSQL_URL -benchmem -cpuprofile profile
goos: windows
goarch: amd64
pkg: go_sql_benchmarks
cpu: AMD Ryzen 7 5800H with Radeon Graphics

Benchmark_sq_sqlite-16               440           2654696 ns/op         1223230 B/op      34837 allocs/op
Benchmark_sqlx_sqlite-16             457           2588915 ns/op          662767 B/op      32697 allocs/op

Benchmark_sq_postgres-16             290          27419349 ns/op          942023 B/op      17426 allocs/op
Benchmark_sqlx_postgres-16           100          20488475 ns/op          383366 B/op      15289 allocs/op

Benchmark_sq_mysql-16                412           3157868 ns/op          987393 B/op      17677 allocs/op
Benchmark_sqlx_mysql-16              788           3045322 ns/op          418760 B/op      13636 allocs/op

PASS
ok      go_sql_benchmarks       18.686s

PS C:\Users\cbw\Documents\go_sql_benchmarks> go test -bench . -postgres $env:POSTGRES_URL -mysql $env:MYSQL_URL -benchmem -cpuprofile profile
goos: windows
goarch: amd64
pkg: go_sql_benchmarks
cpu: AMD Ryzen 7 5800H with Radeon Graphics

Benchmark_sq_sqlite-16               436           2651153 ns/op         1223216 B/op      34837 allocs/op
Benchmark_sqlx_sqlite-16             463           2557336 ns/op          662759 B/op      32697 allocs/op

Benchmark_sq_postgres-16             100          20321208 ns/op          942675 B/op      17428 allocs/op
Benchmark_sqlx_postgres-16           100          23468400 ns/op          383365 B/op      15289 allocs/op

Benchmark_sq_mysql-16                733           5041293 ns/op          987402 B/op      17677 allocs/op
Benchmark_sqlx_mysql-16              787           3870108 ns/op          418757 B/op      13636 allocs/op

PASS
ok      go_sql_benchmarks       15.305s
```
