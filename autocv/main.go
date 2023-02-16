package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Link struct {
	Text string
	Url  string
}

type Data struct {
	Title    string
	Subtitle string
	Date     string
	Text     string
	Style    string
	Link     Link
}

type Section struct {
	Title string
	Data  []Data
}

type TemplateData struct {
	Filename    string
	Filetitle   string
	Title       string
	Information []Link
	Sections    []Section
}

// Reads a file from the given path.
func readFile(path string) ([]byte, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("path '%s' does not exist", path)
		} else {
			return nil, fmt.Errorf("invalid path '%s'", path)
		}
	}
	if fileInfo.IsDir() {
		return nil, fmt.Errorf("path '%s' is a directory", path)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file '%s'", path)
	}
	return data, nil
}

// Write the data to the given path.
func writeOutput(data []byte, path string) (string, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("path '%s' does not exist", path)
		} else {
			return "", fmt.Errorf("invalid path '%s'", path)
		}
	}
	if !fileInfo.IsDir() {
		return "", fmt.Errorf("path '%s' is not a directory", path)
	}

	outputFile := filepath.Join(path, "output.txt")
	if err := ioutil.WriteFile(outputFile, data, 0644); err != nil {
		return "", fmt.Errorf("could not write output to path '%s'", path)
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

	data, err := readFile(*dataPath)
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

	templateData := TemplateData{}
	if err := yaml.Unmarshal(data, &templateData); err != nil {
		fmt.Printf("Unable to read data file '%s'.\n", *dataPath)
		os.Exit(1)
	}
	fmt.Printf("FILENAME: %s, FILETITLE: %s, TITLE: %s\n", templateData.Filename, templateData.Filetitle, templateData.Title)

	outputFile, err := writeOutput(template, *outputPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Successfully wrote to file '%s'.", outputFile)
}
