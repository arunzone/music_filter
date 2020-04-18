package music

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	location := "/Users/arun/Downloads/Music"
	files, err := ioutil.ReadDir(location)
	if err == nil {
		for _, file := range files {
			fmt.Println(AbsolutePath(location, file.Name()))
			err := Unzip(AbsolutePath(location, file.Name()))
			if err != nil {
				panic(err)
			}
		}
	}

}

func AbsolutePath(parent, fileName string) string {
	return filepath.Join(parent, fileName)
}

func Unzip(src string) error {
	zipReader, err := zip.OpenReader(src)
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
	return nil
}
