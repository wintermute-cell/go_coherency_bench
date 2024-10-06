PHONY: all

all:
	@echo "Running all benchmarks"
	perf stat -e cache-references,cache-misses go test -bench BenchmarkDenseAccessReadOnly
	perf stat -e cache-references,cache-misses go test -bench BenchmarkDenseAccessReadWrite
	perf stat -e cache-references,cache-misses go test -bench BenchmarkSparseAccessReadOnly
	perf stat -e cache-references,cache-misses go test -bench BenchmarkSparseAccessReadWrite
	@echo "Done..."

