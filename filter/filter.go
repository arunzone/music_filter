package filter

import (
	"path/filepath"
	"strings"
)

func FilterFlac(fileFullPaths []string) []string{
	var filteredFullPaths []string
	for _, fileFullPath := range fileFullPaths {
		_, fileName := filepath.Split(fileFullPath)
		extension := filepath.Ext(fileName)
		if strings.EqualFold(extension, ".flac") {
			filteredFullPaths = append(filteredFullPaths, fileFullPath)
		}
	}
	return filteredFullPaths
}