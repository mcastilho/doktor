package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func removeIOC(ioc string) {

	fmt.Printf("Removing %s\n", ioc)

	err := os.RemoveAll(ioc)

	if err != nil {
		log.Printf("ERROR: %s", err)
	}

}

func Clean(c *cli.Context) {
	results := performScan()

	if len(results) > 0 {
		printResults(results)
		for _, v := range results {
			for _, path := range v.Paths {
				removeIOC(path)
			}
		}
	} else {
		fmt.Println("No malware found! (っ◕‿◕)っ")
	}
}
