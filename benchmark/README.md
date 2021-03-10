## Optimizing dna_checker benchmarks for change in dna_checker
Original
```goos: linux
goarch: amd64
pkg: github.com/gusgins/meli-backend/utils
cpu: Intel(R) Core(TM) i5-9400F CPU @ 2.90GHz
BenchmarkIsMutant-6   	 1477029	       796.8 ns/op
PASS
ok  	github.com/gusgins/meli-backend/utils	11.096s
```
New Versión (Synchronous)
```goos: linux
goarch: amd64
pkg: github.com/gusgins/meli-backend/utils
cpu: Intel(R) Core(TM) i5-9400F CPU @ 2.90GHz
BenchmarkIsMutant-6       826862              1514 ns/op
PASS
ok      github.com/gusgins/meli-backend/utils   4.310s
```

New Versión (Asynchronous)
```goos: linux
goarch: amd64
pkg: github.com/gusgins/meli-backend/utils
cpu: Intel(R) Core(TM) i5-9400F CPU @ 2.90GHz
BenchmarkIsMutant-6   	  154576	      7824 ns/op
PASS
ok  	github.com/gusgins/meli-backend/utils	1.890s
```
