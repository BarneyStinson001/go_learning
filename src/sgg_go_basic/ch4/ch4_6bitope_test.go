package main

import "testing"

func TestBitOperation(t *testing.T) {
	a := 3//11
	b := 7//111
	t.Log("3&7:",a&b)//11有0则0
	t.Log("3|7:",a|b)//111有1则1

	c:=1//01
	d:=2//10
	t.Log("1|2:",c|d)//11有1则1
	t.Log("1^2:",c^d)//不同为1

	e:=b>>1
	f:=a<<1
	t.Log("7>>1:",e)  //右移n位,除
	t.Log("3<<1",f)   //左移，乘
}