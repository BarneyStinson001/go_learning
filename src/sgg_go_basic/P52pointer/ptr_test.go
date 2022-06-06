package P52pointer

import (
	"testing"
)

func TestPtr(t *testing.T)  {
 var i = 10
 t.Log(&i)

 var p0 *int = &i
 t.Log(p0)
 t.Log(*p0)

 var s = "hello"
 t.Log(&s)

 p:=&s
 t.Log(*p)
}
