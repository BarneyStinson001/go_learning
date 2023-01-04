package main

import "testing"

func TestStringRange(t *testing.T)  {
	var s = "hello,go!"
	for _,v:=range s{
		t.Logf("%c",v)
	}

	for i:=0;i<len(s);i++{
		t.Logf("%c\n",s[i])
	}
}

func TestChineseString(t *testing.T) {
	var str = "您好，go！"
	for i:=0;i<len(str);i++{
		t.Log(str[i])
		t.Logf("%c",str[i])//乱码
	}

	str2:=[]rune(str)
	t.Log(len(str2))
	for i:=0;i<len(str2);i++{
		t.Log(str2[i])
		t.Logf("%c",str2[i])//乱码
	}

	//for idx,val :=range str{
	//	t.Logf("%d: %c",idx,val)
	//}
}