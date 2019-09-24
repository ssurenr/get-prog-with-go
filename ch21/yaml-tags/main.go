package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func main() {
	type location struct {
		Lat  float64 `yaml:"latitude"`
		Long float64 `yaml:"longitude"`
	}

	curiosity := location{-4.5895, 137.4417}

	bytes, err := yaml.Marshal(curiosity)
	exitOnError(err)
	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
