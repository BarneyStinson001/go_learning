package cal_test

import "testing"

func TestCompareArray(t *testing.T)  {
	a:=[...]int{1,2,3,4}
	b:=[...]int{1,2,3,5}
	//c:=[...]int{1,2,3,4,5}
	d:=[...]int{1,2,3,4}
	t.Log(a==b)
	//t.Log(a==c)//不能比较，元素个数不一样
	t.Log(a==d)

}

const  (
	Readable = 1 <<iota
	Writable
	Executable
)
func TestBitClear(t *testing.T) {
	a:=7
	a =a &^ Readable
	a= a &^ Executable
	t.Log(a&Readable==Readable,a&Writable==Writable,a&Executable==Executable)
}