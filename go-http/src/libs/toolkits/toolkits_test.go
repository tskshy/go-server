package toolkits

import (
	"testing"
)

/* 测试某个package
go test -v -count=1 libs/toolkits
*/
/* 测试某个方法
go test libs/toolkits -test.run Test_BinPath -v
*/
func Test_CompareWithN(t *testing.T) {
	var r = CompareWithN(2)(3)

	if r != 3 {
		t.Error("Need result 3")
		return
	}
	t.Log("test ok\n")
}

func Test_BinPath(t *testing.T) {
	var s = BinPath()
	t.Log(s)
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
