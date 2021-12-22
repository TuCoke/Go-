package main

import (
	"fmt"
	"main/Pipline"
)

type City interface {
	FetchData()
}

func main() {
	fmt.Println("MAIN")
	test:=Pipline.NewTestInfo()
	test.FetchData()
}
