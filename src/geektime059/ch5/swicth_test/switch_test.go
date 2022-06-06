package swicth_test

import (
	"runtime"
	"testing"
)

func TestSwitch(t *testing.T) {
	switch os :=runtime.GOOS;os {
	case "linux":
		t.Log("Linux")
	case "Mac":
		t.Log("Mac")
	default:
		t.Logf("%s",os)


	}
}
func TestMultiCases(t *testing.T)  {
	Num:=5
	switch  {
	case 0<=Num && Num<=3:
		t.Log("0-3")
	case 4 <= Num && Num <= 6:
		t.Log("4-6")
	case 7 <= Num && Num <= 9:
		t.Log("7-9")
	}
}