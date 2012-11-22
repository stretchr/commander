package benchmarks

import (
	"testing"
)

type TheObject struct {
	Field string
}

// DoSomething is needed in order to use the object we are creating
func DoSomething(o *TheObject) {}

func Benchmark_CreateObject_UsingNew(b *testing.B) {

	for i := 0; i < b.N; i++ {
		o := new(TheObject)
		o.Field = "something"
		DoSomething(o)
	}

}

func Benchmark_CreateObject_UsingBraces(b *testing.B) {

	for i := 0; i < b.N; i++ {
		DoSomething(&TheObject{"something"})
	}

}
