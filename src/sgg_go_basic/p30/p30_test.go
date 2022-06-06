package p30_test

import "testing"

func TestHello(t *testing.T)  {
	t.Log("Hello,world")
}

func TestPrint(t * testing.T)  {
	t.Log("\nlisi\n12\nbeijing\nshanghai")
}

func TestPrint2(t * testing.T)  {
	str:=`
     *       *
   *    *  *    *
  *       *       *
*                   *
********************** `
t.Log(str)
}