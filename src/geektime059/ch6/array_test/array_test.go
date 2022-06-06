package array_test

import (
	"testing"
)

func TestArratInit(t *testing.T) {
	var arr [3]int

	arr1:=[...]int{1,2,3}
	arr2:=[4]int{1,2,3,4}
	arr1[0]=4

	t.Log(arr,arr1,arr2)
}

func TestArraytravel(t *testing.T) {
	arr3 :=[...]int{2,4,6}
	for i:=0;i<len(arr3);i++{
		t.Log(arr3[i])
	}
	for idx,v :=range arr3{
		t.Log(idx,v)
	}
	for _,v :=range arr3{
		t.Log(v)
	}
}