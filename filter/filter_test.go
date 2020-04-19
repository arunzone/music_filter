package filter

import (
	"reflect"
	"testing"
)

func TestFilterFlac(t *testing.T) {
	fileFullPaths := []string{
		"/a/b/one.txt",
		"/a/b/two.flac"}

	flacFilePaths := FilterFlac(fileFullPaths)

	expectedPaths := []string{"/a/b/two.flac"}

	if !reflect.DeepEqual(flacFilePaths, expectedPaths) {
		t.Errorf("expected '%v' but got '%v'", expectedPaths, flacFilePaths)
	}
}
