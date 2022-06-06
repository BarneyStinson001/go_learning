package main

import "testing"

func TestCal(t *testing.T)  {
	var a,b = +3,-4
	t.Logf("a , b %d\t %d",a,b)

	a=a+b
	t.Logf("tmp_a %d",a)

	b=a-b
	a=a-b
	t.Logf("a , b %d\t %d",a,b)

	c:=a*b
	t.Logf("c %d",c)

	d:=a/b
	t.Logf("d %d",d)

	e:=a%b
	t.Logf("e %d",e)

	a++
	b--
	t.Logf("a , b %d\t %d",a,b)

	s1:="aaa"
	s2:="bbb"
	t.Logf(s1+s2)

}