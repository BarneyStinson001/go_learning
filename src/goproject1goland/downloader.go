package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp,err :=http.Get("http://www.imooc.com")
	if err !=nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes,_ := ioutil.ReadAll(resp.Body)  //这里不容易出错
	fmt.Println(bytes)

}
