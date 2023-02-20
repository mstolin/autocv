package main

import "testing"

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
