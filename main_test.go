package main

import (
	"math/rand"
	"runtime"
	"runtime/debug"
	"testing"

	_ "net/http/pprof"
)

const dataSize = 10000000 // Size of the data array

// generateData creates a slice of random integers.
func generateData() []int {
	data := make([]int, dataSize)
	for i := range data {
		data[i] = rand.Intn(1000)
	}
	return data
}

// generateIndices creates a slice of random indices for scattered access.
func generateIndices() []int {
	indices := make([]int, dataSize)
	for i := range indices {
		indices[i] = rand.Intn(dataSize)
	}
	return indices
}

// BenchmarkDenseAccessReadOnly tests sequential read-only access.
func BenchmarkDenseAccessReadOnly(b *testing.B) {
	data := generateData()
	indices := generateIndices()
	sum := 0

	runtime.GC()                  // Run GC once before the benchmark starts
	debug.SetGCPercent(-1)        // Disable garbage collection
	defer debug.SetGCPercent(100) // Re-enable GC after the benchmark

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < dataSize-2; j++ {
			_ = indices[j] // Access indices for fairness
			sum += data[j] + data[j+1]
		}
	}
	b.StopTimer()
	_ = sum // Prevent sum from being optimized away
}

// BenchmarkDenseAccessReadWrite tests sequential read-write access.
func BenchmarkDenseAccessReadWrite(b *testing.B) {
	data := generateData()
	indices := generateIndices()

	runtime.GC()                  // Run GC once before the benchmark starts
	debug.SetGCPercent(-1)        // Disable garbage collection
	defer debug.SetGCPercent(100) // Re-enable GC after the benchmark

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < dataSize-2; j++ {
			_ = indices[j] // Access indices for fairness
			data[j+1] = data[j] + data[j+1]
		}
	}
	b.StopTimer()
}

// BenchmarkSparseAccessReadOnly tests random read-only access using precomputed indices.
func BenchmarkSparseAccessReadOnly(b *testing.B) {
	data := generateData()
	indices := generateIndices()
	sum := 0

	runtime.GC()                  // Run GC once before the benchmark starts
	debug.SetGCPercent(-1)        // Disable garbage collection
	defer debug.SetGCPercent(100) // Re-enable GC after the benchmark

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < dataSize-2; j++ {
			k := indices[j]
			sum += data[j] + data[k]
		}
	}
	b.StopTimer()
	_ = sum // Prevent sum from being optimized away
}

// BenchmarkSparseAccessReadWrite tests random read-write access using precomputed indices.
func BenchmarkSparseAccessReadWrite(b *testing.B) {
	data := generateData()
	indices := generateIndices()

	runtime.GC()                  // Run GC once before the benchmark starts
	debug.SetGCPercent(-1)        // Disable garbage collection
	defer debug.SetGCPercent(100) // Re-enable GC after the benchmark

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < dataSize-2; j++ {
			k := indices[j]
			data[k] = data[j] + data[k]
		}
	}
	b.StopTimer()
}
