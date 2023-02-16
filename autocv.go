package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Reads a file from the given path.
func readFile(path string) ([]byte, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("Path '%s' does not exist.", path)
		} else {
			return nil, fmt.Errorf("Invalid path '%s'.", path)
		}
	}
	if fileInfo.IsDir() {
		return nil, fmt.Errorf("Path '%s' is a directory.", path)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Could not read file '%s'.", path)
	}
	return data, nil
}

// Write the data to the given path.
func writeOutput(data []byte, path string) (string, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("Path '%s' does not exist.", path)
		} else {
			return "", fmt.Errorf("Invalid path '%s'.", path)
		}
	}
	if !fileInfo.IsDir() {
		return "", fmt.Errorf("Path '%s' is not a directory.", path)
	}

	outputFile := filepath.Join(path, "output.txt")
	if err := ioutil.WriteFile(outputFile, data, 0644); err != nil {
		return "", fmt.Errorf("Error writing output to path '%s'", path)
	}

	return outputFile, nil
}

func main() {
	dataPath := flag.String("data", "", "Path to the data file.")
	templatePath := flag.String("template", "", "Path to the template file.")
	outputPath := flag.String("output", "", "Output path.")
	flag.Parse()

	// all flags are required
	if *dataPath == "" || *templatePath == "" || *outputPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	_, err := readFile(*dataPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Successfully red data from '%s'.\n", *dataPath)

	template, err := readFile(*templatePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Successfully red template from '%s'.\n", *templatePath)

	outputFile, err := writeOutput(template, *outputPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Successfully wrote to file '%s'.", outputFile)
}
