package filter

import (
	"path/filepath"
	"strings"
)

func FilterBy(fileFullPaths []string, extension string) []string{
	var filteredFullPaths []string
	for _, fileFullPath := range fileFullPaths {
		_, fileName := filepath.Split(fileFullPath)
		if strings.EqualFold(filepath.Ext(fileName), extension) {
			filteredFullPaths = append(filteredFullPaths, fileFullPath)
		}
	}
	return filteredFullPaths
}