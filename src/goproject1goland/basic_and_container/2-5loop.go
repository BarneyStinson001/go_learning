package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int	) string  {
	result := ""
	for ;n>0;n=n/2{
		lsb := n%2
		result =strconv.Itoa(lsb) +result
	}
	return result
}
func printFile(filename string)  {
	file,err :=os.Open(filename)
	if err!= nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

}

func forever()  {
	for  {
		fmt.Println("abc")
	}
}


func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13))
	printFile("abc.txt")
	forever()
}
