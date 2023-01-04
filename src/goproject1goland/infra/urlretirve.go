package infra

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct {}

func (Retriever) Get(url string) string{//大写那边才能使用
		resp,err :=http.Get(url)
		if err !=nil {
			panic(err)
		}
		defer resp.Body.Close()

		bytes,_ := ioutil.ReadAll(resp.Body)  //这里不容易出错
		return string(bytes)//也可以string(bytes)
}