package main

import (
	"fmt"
	"io/ioutil"
	"music_filter/music"
	"strings"
)

func main() {
	location := "/Users/arun/Downloads/Music"
	files, err := ioutil.ReadDir(location)
	if err == nil {
		for _, file := range files {
			fullPathFileName := music.AbsolutePath(location, file.Name())
			fmt.Println(fullPathFileName)
			if !IsHiddenFile(file.Name()) {
				err := music.Unzip(fullPathFileName, music.FileNameWithoutExtension(fullPathFileName))
				if err != nil {
					panic(err)
				}
			}
		}
	}

}

func IsHiddenFile(fileName string) bool {
	return strings.HasPrefix(fileName, ".")
}

