package main

import (
	"fmt"
	"io/ioutil"
	"music_filter/music"
)

func main() {
	location := "/Users/arun/Downloads/Music"
	files, err := ioutil.ReadDir(location)
	if err == nil {
		for _, file := range files {
			fullPathFileName := music.AbsolutePath(location, file.Name())
			fmt.Println(fullPathFileName)
			err := music.Unzip(fullPathFileName, music.FileNameWithoutExtension(fullPathFileName))
			if err != nil {
				panic(err)
			}
		}
	}

}

