package p20_test

import "testing"

func TestEscapeChar(t *testing.T)  {
	t.Log("lisi\thello\t!")
	t.Log("lisia\thello\t!")
	t.Log("lisiaa\thello\t!")
	t.Log("lisiaaa\thello\t!")
	t.Log("lisiaaaa\thello\t!")
	t.Log("lisiaaaaa\thello\t!")
	t.Log("lisiaaaaaa\thello\t!")
	t.Log("lisiaaaaaaa\thello\t!")

	t.Log("zhancgasksacasc\rabcdef")
	t.Log("\n姓名\t年龄\t籍贯\t住址\njoln\t12\t河北\t北京")
}