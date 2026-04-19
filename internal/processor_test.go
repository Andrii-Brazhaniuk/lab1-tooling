package internal

import "testing"

var testData = []string{"apple", "123", "banana", "Cherry", "go"}

func BenchmarkProcessDataSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessDataSlow(testData)
	}
}

func BenchmarkProcessDataFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessDataFast(testData)
	}
}