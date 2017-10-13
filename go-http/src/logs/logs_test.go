package logs

import (
	"fmt"
	"testing"
)

func Test_Log(t *testing.T) {
	Debug("hello")
}

func Benchmark_Log(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("")
	}
}
