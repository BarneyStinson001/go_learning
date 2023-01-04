package strcon_test

import (
	"strconv"
	"testing"
)

func TestConvToString(t *testing.T) {
	var i = 0
	var f = 0.68
	var b bool = false

	//var s string

	stri := strconv.FormatInt(int64(i), 10)
	t.Logf("str %q\ttype %T\n", stri, stri)

	strf := strconv.FormatFloat(float64(f),'E',-1,32)
	t.Logf("str %q\ttype %T\n",strf,strf)

	strb := strconv.FormatBool(b)
	t.Log()
	t.Logf("str %q\ttype %T\n",strb,strb)
}

func TestConvFromString(t *testing.T)  {
	var str = "true"
	b,_ := strconv.ParseBool(str)
	t.Logf("str %v\ttype %T\n",b,b)
	var str2 = "5"
	i,_:=strconv.ParseInt(str2,10,32)
	t.Logf("str %v\ttype %T\n",i,i)

	var str3 = "hello"
	i2,_ := strconv.ParseInt(str3,10,64)//返回0
	t.Log(i2)

}