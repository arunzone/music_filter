package music

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAbsolutePath(t *testing.T) {
	absolutePath := AbsolutePath("a", "b")
	expected := "a/b"

	assertEquals(absolutePath, expected, t)
}

func assertEquals(absolutePath string, expected string, t *testing.T) {
	if absolutePath != expected {
		t.Errorf("expected '%v' but got '%v'", expected, absolutePath)
	}
}

func TestUnzipNotReturnError(t *testing.T) {
	getwd, err := os.Getwd()
	if err != nil {
		t.Fatalf(" Problem in finding current working directory")
	}
	absPath, _ := filepath.Abs(getwd + "/../test/test.txt.zip")
	destinationPath, _ := filepath.Abs(getwd + "/../test/test.txt")
	zipStatus := Unzip(absPath, destinationPath)

	if zipStatus != nil {
		t.Errorf("expected '%v' but got '%v'", nil, zipStatus)
	}

	stat, err := os.Stat(destinationPath)
	if err != nil {
		t.Errorf("expected '%v' but got '%v'", true, err)
	}

	if !stat.IsDir() {
		t.Errorf("expected '%v' but got '%v'", true, stat.IsDir())
	}

}

func TestFileNameWithoutExtension(t *testing.T) {
	name := FileNameWithoutExtension("abc.txt")

	assertEquals(name, "abc", t)
}
