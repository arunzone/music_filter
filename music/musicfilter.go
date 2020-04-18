package music

import (
	"archive/zip"
	"fmt"
	"io"
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

func Unzip(fileFullPath, destinationLocation string) error {
	zipReader, err := zip.OpenReader(fileFullPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := zipReader.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(destinationLocation, 0755)

	extractAndWriteFile := func(file *zip.File) error {
		fileHandle, err := file.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := fileHandle.Close(); err != nil {
				panic(err)
			}
		}()

		destinationFilePath := filepath.Join(destinationLocation, file.Name)

		os.MkdirAll(filepath.Dir(destinationFilePath), file.Mode())
		destinationFileHandle, err := os.OpenFile(destinationFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer func() {
			if err := destinationFileHandle.Close(); err != nil {
				panic(err)
			}
		}()

		_, err = io.Copy(destinationFileHandle, fileHandle)
		if err != nil {
			return err
		}
		return nil
	}

	for _, file := range zipReader.File {
		fmt.Println(file.Name)
		extractAndWriteFile(file)
	}

	return nil
}
