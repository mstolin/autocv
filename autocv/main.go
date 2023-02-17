package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

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
	Text     StringArray
	Layout   string
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

// See https://github.com/go-yaml/yaml/issues/100#issuecomment-901604971
type StringArray []string

func (a *StringArray) UnmarshalYAML(value *yaml.Node) error {
	var multi []string
	err := value.Decode(&multi)
	if err != nil {
		var single string
		err := value.Decode(&single)
		if err != nil {
			return err
		}
		*a = []string{single}
	} else {
		*a = multi
	}
	return nil
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

/// Renders the given data into the template
func renderTemplate(templatePath, destination string, data TemplateData) error {
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	if err := template.Execute(outputFile, data); err != nil {
		return err
	}
	outputFile.Close()
	return nil
}

/// Generates the destination file for the tex document
func genDestinationPath(destDir, filename string) (string, error) {
	fileInfo, err := os.Stat(destDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("path '%s' does not exist", destDir)
		} else {
			return "", fmt.Errorf("invalid path '%s'", destDir)
		}
	}
	if !fileInfo.IsDir() {
		return "", fmt.Errorf("path '%s' is not a directory", destDir)
	}

	return filepath.Join(destDir, fmt.Sprintf("%s.tex", filename)), err
}

func main() {
	// required CLI flags
	dataPath := flag.String("data", "", "Path to the data file.")
	templatePath := flag.String("template", "", "Path to the template file.")
	outputPath := flag.String("output", "", "Output path.")
	flag.Parse()

	// all flags are required
	if *dataPath == "" || *templatePath == "" || *outputPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// read YAML data
	data, err := readFile(*dataPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmarshal yaml data
	templateData := TemplateData{}
	if err := yaml.Unmarshal(data, &templateData); err != nil {
		fmt.Printf("Unable to read data file '%s'.\n", *dataPath)
		os.Exit(1)
	}
	if templateData.Filename == "" {
		fmt.Println("The property filename must be set.")
		os.Exit(1)
	}

	// Parse template
	destination, err := genDestinationPath(*outputPath, templateData.Filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Write latex document
	if err := renderTemplate(*templatePath, destination, templateData); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
