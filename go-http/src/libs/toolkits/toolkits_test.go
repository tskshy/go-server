package toolkits

import (
	"testing"
)

// go test -v -count=1 libs/toolkits
func Test_CompareWithN(t *testing.T) {
	var r = CompareWithN(2)(3)

	if r != 3 {
		t.Error("Need result 3")
		return
	}
	t.Log("test ok\n")
}

// go test -bench="Benchmark_CompareWithN" -count=5 libs/toolkits
func Benchmark_CompareWithN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var r = CompareWithN(2)(3)
		if r != 3 {
			b.Error("hehe")
		}
	}
}
