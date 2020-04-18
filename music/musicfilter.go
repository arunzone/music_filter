package music

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

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

	os.MkdirAll(destinationLocation, 0777)

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

		if file.FileInfo().IsDir() {
			os.MkdirAll(destinationFilePath, 0777)
		} else {
			os.MkdirAll(filepath.Dir(destinationFilePath), 0777)
			destinationFileHandle, err := os.OpenFile(destinationFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
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
		}
		return nil
	}

	for _, file := range zipReader.File {
		fmt.Println(file.Name)
		extractAndWriteFile(file)
	}

	return nil
}
