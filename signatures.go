package main

import (
	"gopkg.in/yaml.v2"
	"log"
)

func loadSignatures() []IOC {

	t := make(map[string][]string)

	err := yaml.Unmarshal([]byte(signatures), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	loadedSignatures := make([]IOC, 0)

	// load raw yaml into organized structs, probably a better way to do this
	for k, v := range t {
		loadedSignatures = append(loadedSignatures, IOC{Name: k, Paths: v})
	}

	return loadedSignatures
}
