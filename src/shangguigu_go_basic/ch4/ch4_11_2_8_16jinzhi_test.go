package main

import "testing"

func TestIntChange(t *testing.T) {
	var i,j,k,h,l,m int = 5,10,20,40,80,160
	t.Logf("%b\n%b\n%b\n%b\n%b\n%b\n",i,j,k,h,l,m)


	var eight=010
	t.Log(eight)

	var sixteen=0x10
	t.Log(sixteen)
}