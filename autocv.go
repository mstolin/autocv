package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type Link struct {
	Text string
	Url  string
}

type Data struct {
	Title    string
	Subtitle string
	Location string
	Date     string
	Text     []string
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

const HELP_INFO = "Latex document parser\n\n" +
	"USAGE:\n" +
	"  autocv [OPTIONS] [CONFIG]...\n\n" +
	"OPTIONS:\n"

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

/// Simply a - b
func minus(a, b int) int {
	return a - b
}

/// Returns custom functions for the template
func getCustomFuncs() template.FuncMap {
	return template.FuncMap{
		"minus": minus,
	}
}

/// Renders the given data into the template
func renderTemplate(templatePath, destination string, data TemplateData) error {
	content, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return err
	}
	customFuncs := getCustomFuncs()
	tmpl, err := template.New("resume").Funcs(customFuncs).Parse(string(content))
	if err != nil {
		return err
	}
	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	if err := tmpl.Execute(outputFile, data); err != nil {
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

/// Returns only the name of the given config file
func splitFilename(configFile string) string {
	return configFile[:len(configFile)-len(filepath.Ext(configFile))]
}

func main() {
	// required CLI flags
	templatePath := flag.String("template", "", "The Latex template.")
	outputPath := flag.String("outputDir", ".", "Output directory.")
	help := flag.Bool("help", false, "Print help information")
	flag.Parse()
	configs := flag.Args()

	if *help {
		fmt.Print(HELP_INFO)
		flag.PrintDefaults()
		os.Exit(0)
	}

	// at leas one data file must be given
	if len(configs) < 1 {
		fmt.Println("At least one config file must be given.")
		os.Exit(1)
	}

	// template is requires
	if *templatePath == "" {
		fmt.Println("No template given.")
		os.Exit(1)
	}

	for _, configPath := range configs {
		// read json
		data, err := readFile(configPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Unmarshal json data
		var templateData TemplateData
		if err := json.Unmarshal(data, &templateData); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		filename := splitFilename(configPath)

		// Parse template name
		destination, err := genDestinationPath(*outputPath, filename)
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
}
