package main

import (
	"bytes"
	"testing"
)

func TestSplitFilename(t *testing.T) {
	// test empty
	_, err := splitFilename("")
	if err == nil {
		t.Errorf("FAIL splitFilename: There has to be an error for an empty configFile path")
	}

	// test with filename without extension
	path := "../test/testfile"
	got, err := splitFilename(path)
	if err != nil {
		t.Errorf("FAIL splitFilename: There can't be an error for path '%s'", path)
	}
	if got != "testfile" {
		t.Errorf("FAIL splitFilename: Got '%s', expected 'testfile' with path '%s'", got, path)
	}

	// test with a correct path
	path = "../test/testfile.json"
	got, err = splitFilename(path)
	if err != nil {
		t.Errorf("FAIL splitFilename: There can't be an error for path '%s'", path)
	}
	if got != "testfile" {
		t.Errorf("FAIL splitFilename: Got '%s', expected 'testfile' with path '%s'", got, path)
	}
}

func TestMinus(t *testing.T) {
	// test normal ints
	expect := 4
	got := minus(8, 4)
	if expect != got {
		t.Errorf("FAIL minus: 8 - 4 is %d, got %d", expect, got)
	}

	// test with 0
	expect = -8
	got = minus(0, 8)
	if expect != got {
		t.Errorf("FAIL minus: 0 - 8 is %d, got %d", expect, got)
	}
}

func TestReadFile(t *testing.T) {
	// test no existing path
	_, err := readFile("i/do/not.exist")
	if err == nil {
		t.Errorf("FAIL readFile: There must be an error for an empty path")
	}

	// test directory
	_, err = readFile("/")
	if err == nil {
		t.Errorf("FAIL readFile: There must be an error a directory path")
	}

	// test existing file
	content, err := readFile("./LICENSE")
	if err != nil {
		t.Errorf("FAIL readFile: error must be nil for existing file")
	}
	contentLen := len(string(content))
	if contentLen == 0 {
		t.Errorf("FAIL readFile: len of LICENSE file can't be empty")
	}
}

func TestRenderTemplate(t *testing.T) {
	tmpl, err := renderTemplate("{{ .Title }}={{ minus 8 4 }}")
	if err != nil {
		t.Errorf("FAIL renderTemplate: %e", err)
	}
	var data TemplateData
	data.Title = "Test"

	var out bytes.Buffer
	if err = tmpl.Execute(&out, data); err != nil {
		t.Errorf("FAIL renderTemplate: %e", err)
	}
	got := out.String()
	expect := "Test=4"
	if got != expect {
		t.Errorf("FAIL renderTemplate: Expected: '%s', got: '%s'", expect, got)
	}
}

func TestGenDestinationPath(t *testing.T) {
	// test with non existing path
	_, err := genDestinationPath("i/do/not/exist", "file")
	if err == nil {
		t.Errorf("FAIL genDestinationPath: Error can't be nil for non-existing path")
	}

	// test with file as destination
	_, err = genDestinationPath("./LICENSE", "file")
	if err == nil {
		t.Errorf("FAIL genDestinationPath: There must be an error if destination is a file")
	}

	// test with empty strings
	_, err = genDestinationPath("", "")
	if err == nil {
		t.Errorf("FAIL genDestinationPath: Expected error for empty arguments")
	}

	// test with valid arguments
	got, err := genDestinationPath(".", "file")
	expected := "file.tex"
	if err != nil {
		t.Errorf("FAIL genDestinationPath: %e", err)
	}
	if expected != got {
		t.Errorf("FAIL genDestinationPath: Expected '%s' got '%s'", expected, got)
	}
}
