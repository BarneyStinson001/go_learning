package main

import (
	"fmt"
	"goproject1goland/retriever/mock"
	real2 "goproject1goland/retriever/real"
)

//Retriever
type Retriever interface {
	Get( url string) string
}

func download(r Retriever)  string {
	return r.Get("https://www.imooc.com")
}
func main() {
	var r Retriever
	r=mock.Retriever{"this is a fake mock data"}
	r=real2.Retriever{}
	fmt.Println(download(r))
}
