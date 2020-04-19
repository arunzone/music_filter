package file

import (
	"io"
	"os"
	"path/filepath"
)

func Copy(files []string, location string) error {
	os.MkdirAll(location, 0777)
	for _, file := range files {
		fileHandle, err := os.Open(file)
		if err != nil {
			return err
		}
		defer func() {
			if err := fileHandle.Close(); err != nil {
				panic(err)
			}
		}()

		_, fileName := filepath.Split(file)
		targetFile := filepath.Join(location, fileName)
		destinationFileHandle, err := os.OpenFile(targetFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
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
