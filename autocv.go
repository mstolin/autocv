package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	dataPath := flag.String("data", "", "Path to the data file.")
	templatePath := flag.String("template", "", "Path to the template file.")
	flag.Parse()

	// data and template flags are required
	if *dataPath == "" || *templatePath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("data: %s, template: %s\n", *dataPath, *templatePath)
}
