package type_test

import "testing"

type MyInt int64
func TestImplicit(t *testing.T) {
	var  a int32= 1
	var b int64
	//b=a
	b=int64(a)

	t.Log(a,b)

	var c MyInt
	//c=int64(a)
c=MyInt(a)
	t.Log(a,b,c)


}

func TestPoint(t *testing.T) {
	a:=1
	pa :=&a
	t.Log(a,pa)
	t.Logf("%T %T",a,pa)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*"+s+"*",len(s))
}