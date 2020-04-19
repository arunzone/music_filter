package unzip

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

func ListFileNamesInZip(fileFullPath, destinationLocation string) []string {
	zipReader, err := zip.OpenReader(fileFullPath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := zipReader.Close(); err != nil {
			panic(err)
		}
	}()

	actualTargetLocation := actualTargetLocation(fileFullPath, destinationLocation)

	fileNames := make([]string, len(zipReader.File), len(zipReader.File))
	for index, file := range zipReader.File {
		destinationFilePath := filepath.Join(actualTargetLocation, file.Name)
		fileNames[index] = destinationFilePath
	}
	return fileNames
}

func actualTargetLocation(fileFullPath string, destinationLocation string) string {
	if destinationLocation == "" {
		return filepath.Dir(fileFullPath)
	}
	return destinationLocation
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
	actualTargetLocation := actualTargetLocation(fileFullPath, destinationLocation)
	os.MkdirAll(actualTargetLocation, 0777)

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

		destinationFilePath := filepath.Join(actualTargetLocation, file.Name)

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
