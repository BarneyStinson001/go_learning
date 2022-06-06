package fib

import (
	"testing"
)
//TestFibList  如果是TestfibList则不能识别为测试用例
func TestFibList(t *testing.T) {
	var a = 1
	var b = 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log()
}

