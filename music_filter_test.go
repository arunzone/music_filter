package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestIsHiddenFile(t *testing.T) {
	isHiddenFile := IsHiddenFile(".DS_Store")

	if !isHiddenFile {
		t.Errorf("expected '%v' but got '%v'", true, isHiddenFile)
	}
}

func TestListFilesOn(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf(" Problem in finding current working directory")
	}
	absPath, _ := filepath.Abs(filepath.Join(pwd,"test", "dir"))

	fileNames := ListFileNamesOn(absPath)

	expectedFileNames := []string{
		absPath+"/test-simple.zip",
		absPath+"/test.csv",
	}

	if !reflect.DeepEqual(fileNames, expectedFileNames) {
		t.Errorf("expected '%v' but got '%v'", expectedFileNames, fileNames)
	}
}