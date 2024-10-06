# Benchmarking sequential vs random data access for CPU cache effectiveness and performance

Deps:

```
go
perf
```

How to run benchmark:

```bash
make all
```

Exemplary result:

```
perf stat -e cache-references,cache-misses go test -bench BenchmarkDenseAccessReadOnly
goos: linux
goarch: amd64
pkg: go_cache_bench
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkDenseAccessReadOnly-8   	     228	   5245528 ns/op
PASS
ok  	go_cache_bench	2.725s

 Performance counter stats for 'go test -bench BenchmarkDenseAccessReadOnly':

        55.384.535      cache-references:u                                                    
        22.779.399      cache-misses:u                   #   41,13% of all cache refs         

       2,993211760 seconds time elapsed

       3,047863000 seconds user
       0,228955000 seconds sys


perf stat -e cache-references,cache-misses go test -bench BenchmarkDenseAccessReadWrite
goos: linux
goarch: amd64
pkg: go_cache_bench
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkDenseAccessReadWrite-8   	      73	  16729779 ns/op
PASS
ok  	go_cache_bench	1.894s

 Performance counter stats for 'go test -bench BenchmarkDenseAccessReadWrite':

       248.385.123      cache-references:u                                                    
       157.661.168      cache-misses:u                   #   63,47% of all cache refs         

       2,155429340 seconds time elapsed

       2,221396000 seconds user
       0,203248000 seconds sys


perf stat -e cache-references,cache-misses go test -bench BenchmarkSparseAccessReadOnly
goos: linux
goarch: amd64
pkg: go_cache_bench
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkSparseAccessReadOnly-8   	     148	   8751764 ns/op
PASS
ok  	go_cache_bench	3.049s

 Performance counter stats for 'go test -bench BenchmarkSparseAccessReadOnly':

       669.115.809      cache-references:u                                                    
       437.648.743      cache-misses:u                   #   65,41% of all cache refs         

       3,301104090 seconds time elapsed

       3,386521000 seconds user
       0,207170000 seconds sys


perf stat -e cache-references,cache-misses go test -bench BenchmarkSparseAccessReadWrite
goos: linux
goarch: amd64
pkg: go_cache_bench
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkSparseAccessReadWrite-8   	       6	 292540034 ns/op
PASS
ok  	go_cache_bench	2.649s

 Performance counter stats for 'go test -bench BenchmarkSparseAccessReadWrite':

       288.344.206      cache-references:u                                                    
       187.258.831      cache-misses:u                   #   64,94% of all cache refs         

       2,900026454 seconds time elapsed

       2,961786000 seconds user
       0,216151000 seconds sys
```

## Open questions
- Are these benchmarks clean or are other operations polluting the results?
- Why are cache-misses so bad even with BenchmarkDenseAccessReadOnly?
- Why does writing degrade cache performance?
- Why is BenchmarkSparseAccessReadWrite so extremely slow? 
