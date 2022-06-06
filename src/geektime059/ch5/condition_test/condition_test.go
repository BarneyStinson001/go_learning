package condition_test

import "testing"

func TestCondition(t *testing.T) {
	var score =61

	if score<60{
		t.Log("fail")
	}else {
		t.Log("pass")
	}
}

func TestIfMultiSec(t *testing.T) {
	if a:= 1==1;a{
		t.Log(a)
	}
}