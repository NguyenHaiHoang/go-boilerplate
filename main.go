package main

import (
	"apus-sample/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil{
		panic(err)
	}
}
