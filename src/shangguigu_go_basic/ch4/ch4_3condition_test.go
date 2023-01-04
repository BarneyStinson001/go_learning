package main

import "testing"

func TestCondition(t *testing.T) {
	t.Log(4==3)
	t.Log(4!=3)
	t.Log(4<3)
	t.Log(4>3)

	t.Log(true&&false)//短路与  第一个false ： OK
	t.Log(true||false)//短路或 第一个true：OK
	t.Log(!true)
}
