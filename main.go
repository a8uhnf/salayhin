package main

import (
	"fmt"

	"github.com/a8uhnf/salyhin/cmd"
)

func main() {
	fmt.Println("hello salayhin!!!!")
	c := cmd.RootCMD()
	if err := c.Execute(); err != nil {
		panic(err)
	}
}
