package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	hr = "--------------------------------------------------------------------"
)

type IOC struct {
	Name  string
	Paths []string
}

type Result struct {
	Name  string
	Paths []string
}

func printResults(results []Result) {
	fmt.Printf("Found %d malware hit(s)! ( ︶︿︶)\n", len(results))
	fmt.Println(hr)
	fmt.Printf("Signature Name\t\tPath\n")
	fmt.Println(hr)
	for _, v := range results {
		for _, path := range v.Paths {
			fmt.Printf("%s\t%s\n", v.Name, path)
		}
	}
	fmt.Println(hr)
}

// where the scanning occurs
func performScan() []Result {

	// load the yaml data from signatures
	loadedSignatures := loadSignatures()

	// empty array used for storing results if we have any
	results := make([]Result, 0)

	// do the scan
	for _, v := range loadedSignatures {

		// Grab current user so we can create an absolute path later
		currentUser, _ := user.Current()

		// create empty array for paths we find
		foundPaths := make([]string, 0)

		// loop through loaded signature paths
		for _, line := range v.Paths {

			// replace %% sqlite wildcard chars with the current username
			iocPath := strings.Replace(line, "%%", currentUser.Username, -1)

			// test to see if this is a wildcard expression and expand it if needed
			isWildcard, _ := regexp.Match("%\\.|%$", []byte(iocPath))

			// if it is a wildcard, do the wild card things
			if isWildcard == true {

				// replace sqlite wildcard with glob wildcard
				wildcardPath := strings.Replace(line, "%", "*", -1)

				// grab any matches if we have em
				wildcardMatches, _ := filepath.Glob(wildcardPath)

				// loop through wildcard matches
				for _, match := range wildcardMatches {

					// test to see if they exist, if they do, add them to the paths we've found
					if _, err := os.Stat(match); err == nil {
						foundPaths = append(foundPaths, match)
					}

				}

			} else {

				// test if the path exists
				if _, err := os.Stat(iocPath); err == nil {
					foundPaths = append(foundPaths, iocPath)
				}
			}
		}

		// only add the paths we've found to the results array if we have em
		if len(foundPaths) > 0 {
			result := Result{Name: v.Name, Paths: foundPaths}
			results = append(results, result)
		}

	}

	return results
}

func Scan(c *cli.Context) {

	results := performScan()
	if !c.Bool("quiet") {
		if len(results) > 0 {
			printResults(results)
		} else {
			fmt.Println("No malware found! (っ◕‿◕)っ")
		}
	}

}
