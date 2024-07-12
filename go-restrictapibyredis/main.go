package main

import (
	"demo/go-restrictapibyredis/service"
	"fmt"
)

func main() {
	err := service.Ask()
	if err != nil {
		fmt.Print(err.Error())
	}

}
