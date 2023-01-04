package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s:= "12345abc一二三四"
	fmt.Println(len(s))
	fmt.Printf("%X\n",[]byte(s))
	for _,b := range []byte(s){
		fmt.Printf("%x  ",b)
	}
	fmt.Printf("\n")
	//utf 可变长，英文一个字节，中文三个字节
	for i,ch := range s{
		fmt.Printf("(%d %X)",i,ch)
	}//ch int32 就是rune
	fmt.Printf("\n")

	fmt.Println("Rune count : ",utf8.RuneCountInString(s))

	byts :=[]byte(s)
	for len(byts)>0 {
		ch, size := utf8.DecodeRune(byts)
		byts = byts[size:]
		fmt.Printf("%c ",ch)
	}

	for i,ch :=range []rune(s){
		fmt.Printf("(%d %c) ",i,ch)
	}


}
