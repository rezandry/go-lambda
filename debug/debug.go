package main

import (
	"fmt"

	"github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

func main() {

	response, err := golambdainvoke.Run(golambdainvoke.Input{
		Port:    8001,
		Payload: "Reza",
	})
	fmt.Println(response)
	s := string(response)
	fmt.Println(s)
	fmt.Println(err)
}
