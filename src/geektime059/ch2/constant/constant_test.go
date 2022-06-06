package constant_test

import "testing"

const (
	Monday = iota +1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func TestConstantIota(t *testing.T) {
	t.Log(Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday)
}

const  (
	Readable = 1 <<iota
	Writavle
	Executable
)
func TestContant(t *testing.T) {
	a:=1
	t.Log(a&Readable==Readable,a&Writavle==Writavle,a&Executable==Executable)
}