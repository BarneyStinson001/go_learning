package main

import "testing"

func TestPrintTenTime(t * testing.T)  {
	for i:=0;i<10;i++{
		t.Log(i,"print something repeatedly")
	}
t.Log("************************")
	//while
	j:=0
	for j<10{
		t.Log(j,"print something repeatedly")
		j++
	}
	t.Log("************************")

	//ultimateloop
	k:=0
	for{
		if k<10 {
			t.Log(k,"print something repeatedly")
		}		else{
			break
		}
		k++
	}



}