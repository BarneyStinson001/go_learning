package main

import "testing"

func TestCount9(t *testing.T)  {
	var max = 100
	base:= 9
	cnt := 0
	for i :=1;i<=max;i++{
		if i%base==0{
			t.Log(i)
			cnt++
		}
	}
	t.Log("total num:",cnt)
}