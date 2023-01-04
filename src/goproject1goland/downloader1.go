package main

import (
	"Project02/infra"
	"fmt"
	"io/ioutil"
	"net/http"
)

func retrive(url string) []byte  {
	resp,err :=http.Get(url)
	if err !=nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes,_ := ioutil.ReadAll(resp.Body)  //这里不容易出错
	return bytes//也可以string(bytes)  直接返回string  ,不用printf()
}
func getRetriever() infra.Retriever  {
	return infra.Retriever{}
}

//?  something that can get url
type retriever interface {
	Get(string) string
}

func main() {
	fmt.Println(retrive("https://www.imooc.com"))
	fmt.Printf("%s \n",retrive("https://www.imooc.com"))
	retriever :=infra.Retriever{}
	fmt.Print(retriever.Get("https://www.imooc.com"))

	//retriever1 :=getRetriever()  等价于下面
	var retriever1 infra.Retriever = getRetriever()
	fmt.Print(retriever1.Get("https://www.imooc.com"))

	//如果要换testing.retriever,需要改好几处。
	//infra 和 testing 的retriever其实都是geturl  功能
	//var retriever1 ？ = getRetriever()



}