package benchmarks

import (
	"testing"
)

const SizeOfArray int = 1000000

func BenchmarkBuildArrayWithAppend(b *testing.B) {

	for i := 0; i < b.N; i++ {

		var bigArray []int
		for count := 0; count < SizeOfArray; count++ {
			bigArray = append(bigArray, count)
		}

	}

}

func BenchmarkBuildArrayWithPredefinedSize(b *testing.B) {

	for i := 0; i < b.N; i++ {

		var bigArray []int = make([]int, SizeOfArray)
		for count := 0; count < SizeOfArray; count++ {
			bigArray[count] = count
		}

	}

}
