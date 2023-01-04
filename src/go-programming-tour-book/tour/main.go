package main

import (
	"github.com/BarneyStinson001/go_project_tour/cmd"
	"log"
)

func main() {
	err:=cmd.Execute()
	if err!=nil{
		log.Fatalf("cmd.Execute err: %v",err)
}
}
