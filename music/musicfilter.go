package music

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	location := "/Users/arun/Downloads/Music"
	files, err := ioutil.ReadDir(location)
	if err == nil {
		for _, file := range files {
			fullPathFileName := AbsolutePath(location, file.Name())
			fmt.Println(fullPathFileName)
			err := Unzip(fullPathFileName, FileNameWithoutExtension(fullPathFileName))
			if err != nil {
				panic(err)
			}
		}
	}

}

func AbsolutePath(parent, fileName string) string {
	return filepath.Join(parent, fileName)
}

func FileNameWithoutExtension(fileName string) string {
	extension := filepath.Ext(fileName)
	return strings.TrimSuffix(fileName, extension)
}

func Unzip(fileFullPath, detinationLocation string) error {
	zipReader, err := zip.OpenReader(fileFullPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := zipReader.Close(); err != nil {
			panic(err)
		}
	}()

	for _, file := range zipReader.File {
		fmt.Println(file.Name)
	}
	os.MkdirAll(detinationLocation, 0755)
	return nil
}
