package main

import (
	"io/ioutil"
	"music_filter/filter"
	"music_filter/unzip"
	"path/filepath"
	"strings"
)

func main() {
	location := "/Users/arun/Downloads/Music"
	destinationLocation := "/Users/arun/Documents/Music"
	fileNames := ListFileNamesOn(location)
	flacFileNames := filter.FilterFlac(fileNames)
	ExtractZipOn(location)

}

func ListFileNamesOn(location string) []string {
	files, err := ioutil.ReadDir(location)
	fileNames := make([]string, len(files), len(files))
	if err == nil {
		for index, file := range files {
			fileNames[index] = unzip.AbsolutePath(location, file.Name())
		}
	}
	return fileNames
}

func ExtractZipOn(location string) {
	files := ListFileNamesOn(location)
	for _, file := range files {
		_, fileNameAlone := filepath.Split(file)
		if !IsHiddenFile(fileNameAlone) {
			err := unzip.Unzip(file, unzip.FileNameWithoutExtension(file))
			if err != nil {
				panic(err)
			}
		}
	}
}

func IsHiddenFile(fileName string) bool {
	return strings.HasPrefix(fileName, ".")
}