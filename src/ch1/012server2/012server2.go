package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int
func main() {
	http.HandleFunc("/",handler)
	http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8001",nil))
}

func handler(w http.ResponseWriter,r *http.Request)  {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
}
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w,"count %d\n",count)
	mu.Unlock()
}

//cmd运行后。 ./009fetch http://localhost:8001/1/2
//cmd运行后。 ./009fetch http://localhost:8001/count
//不要用浏览器，浏览器会自动请求到别的东西，导致count一直在累加