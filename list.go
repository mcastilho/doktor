package main

import (
	"fmt"
	"github.com/urfave/cli"
)

func List(c *cli.Context) {
	loadedSignatures := loadSignatures()
	fmt.Println("Signature Name")
	fmt.Println("--------------")
	for _, v := range loadedSignatures {
		fmt.Printf("%s\n", v.Name)
	}
}
