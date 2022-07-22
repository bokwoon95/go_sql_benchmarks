To run these benchmarks, you need a `$POSTGRES_URL` and `$MYSQL_URL`.

These are the back-to-back runs of the same benchmarks. The results can vary wildly between runs, so even the output below is not necessarily what you will see.


```text
PS C:\Users\cbw\Documents\go_sql_benchmarks> go test -bench . -postgres $env:POSTGRES_URL -mysql $env:MYSQL_URL -benchmem -cpuprofile profile
goos: windows
goarch: amd64
pkg: go_sql_benchmarks
cpu: AMD Ryzen 7 5800H with Radeon Graphics

Benchmark_sq_sqlite-16                       451           2661208 ns/op         1223191 B/op      34837 allocs/op
Benchmark_sq_compiled_sqlite-16              452           2677675 ns/op         1222180 B/op      34831 allocs/op
Benchmark_sqlx_sqlite-16                     464           2573864 ns/op          662750 B/op      32697 allocs/op

Benchmark_sq_postgres-16                     598          23112828 ns/op          941865 B/op      17425 allocs/op
Benchmark_sq_compiled_postgres-16            100          35939766 ns/op          941751 B/op      17423 allocs/op
Benchmark_sqlx_postgres-16                   100          20157269 ns/op          383365 B/op      15289 allocs/op

Benchmark_sq_mysql-16                        356           4257722 ns/op          987379 B/op      17677 allocs/op
Benchmark_sq_compiled_mysql-16               100          11004194 ns/op          986586 B/op      17671 allocs/op
Benchmark_sqlx_mysql-16                      408           2959993 ns/op          419095 B/op      13636 allocs/op

PASS
ok      go_sql_benchmarks       31.690s
```

```text
PS C:\Users\cbw\Documents\go_sql_benchmarks> go test -bench . -postgres $env:POSTGRES_URL -mysql $env:MYSQL_URL -benchmem -cpuprofile profile
goos: windows
goarch: amd64
pkg: go_sql_benchmarks
cpu: AMD Ryzen 7 5800H with Radeon Graphics

Benchmark_sq_sqlite-16                       442           2672059 ns/op         1223203 B/op      34837 allocs/op
Benchmark_sq_compiled_sqlite-16              447           2688661 ns/op         1222182 B/op      34831 allocs/op
Benchmark_sqlx_sqlite-16                     458           2585265 ns/op          662758 B/op      32697 allocs/op

Benchmark_sq_postgres-16                     100          36079354 ns/op          942705 B/op      17429 allocs/op
Benchmark_sq_compiled_postgres-16            100          26839153 ns/op          941833 B/op      17423 allocs/op
Benchmark_sqlx_postgres-16                   100          20143237 ns/op          383364 B/op      15289 allocs/op

Benchmark_sq_mysql-16                        254           6467254 ns/op          987393 B/op      17677 allocs/op
Benchmark_sq_compiled_mysql-16               249           4085895 ns/op          986447 B/op      17671 allocs/op
Benchmark_sqlx_mysql-16                      439           3583004 ns/op          419082 B/op      13636 allocs/op

PASS
ok      go_sql_benchmarks       19.575s
```

```text
PS C:\Users\cbw\Documents\go_sql_benchmarks> go test -bench . -postgres $env:POSTGRES_URL -mysql $env:MYSQL_URL -benchmem -cpuprofile profile
goos: windows
goarch: amd64
pkg: go_sql_benchmarks
cpu: AMD Ryzen 7 5800H with Radeon Graphics

Benchmark_sq_sqlite-16                       442           2678971 ns/op         1223204 B/op      34837 allocs/op
Benchmark_sq_compiled_sqlite-16              427           2639534 ns/op         1222181 B/op      34831 allocs/op
Benchmark_sqlx_sqlite-16                     465           2586066 ns/op          662750 B/op      32697 allocs/op

Benchmark_sq_postgres-16                      99          33157465 ns/op          942781 B/op      17429 allocs/op
Benchmark_sq_compiled_postgres-16            100          32904659 ns/op          941759 B/op      17423 allocs/op
Benchmark_sqlx_postgres-16                   100          26659343 ns/op          383389 B/op      15289 allocs/op

Benchmark_sq_mysql-16                        735           8440821 ns/op          987372 B/op      17677 allocs/op
Benchmark_sq_compiled_mysql-16               250           5301872 ns/op          986480 B/op      17671 allocs/op
Benchmark_sqlx_mysql-16                      807           4541298 ns/op          418775 B/op      13636 allocs/op

PASS
ok      go_sql_benchmarks       25.849s
```
