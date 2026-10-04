package main

import "testing"

func BenchmarkCountNumberOfLines(b *testing.B) {
	for b.Loop() {
		CountNumberOfLines("test.txt")
	}
}
