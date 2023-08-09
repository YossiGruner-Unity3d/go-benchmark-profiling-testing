package main

import (
	"strings"
	"testing"
)

// Benchmark strings.Join
func BenchmarkJoin(b *testing.B) {
	strs := []string{"Hello", "World", "Go", "Benchmarks"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Join(strs, " ")
	}
}

// Benchmark simple concatenation
func BenchmarkConcatenate(b *testing.B) {
	strs := []string{"Hello", "World", "Go", "Benchmarks"}

	for i := 0; i < b.N; i++ {
		var result string
		for _, str := range strs {
			result += " " + str
		}
	}
}
