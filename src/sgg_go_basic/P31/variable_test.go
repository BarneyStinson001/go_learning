package variable_test

import "testing"

func TestVariable(t *testing.T\) {
	var i int  //声明类型，初始化为类型空值
	t.Log(i)  //
	i=10
	t.Log(i)
	var j = 11  //类型推导
	t.Log(j)

	k:=12
	t.Log(k)

	var(
		a=1
		b=2
		c=3
	)
	t.Log(a,b,c)


}
