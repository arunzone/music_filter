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
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf(" Problem in finding current working directory")
	}
	absPath, _ := filepath.Abs(pwd + "/../test/test-simple.zip")
	destinationPath, _ := filepath.Abs(pwd + "/../test/test.txt")
	zipStatus := Unzip(absPath, destinationPath)

	if zipStatus != nil {
		t.Errorf("expected '%v' but got '%v'", nil, zipStatus)
	}

}

func TestUnzipCreateDestinationDirectory(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf(" Problem in finding current working directory")
	}
	absPath, _ := filepath.Abs(pwd + "/../test/test-simple.zip")
	destinationPath, _ := filepath.Abs(pwd + "/../test/test-simple")

	Unzip(absPath, destinationPath)

	stat, err := os.Stat(destinationPath)
	if err != nil {
		t.Errorf("expected '%v' but got '%v'", true, err)
	}

	if !stat.IsDir() {
		t.Errorf("expected '%v' but got '%v'", true, stat.IsDir())
	}
}

func TestUnzipExtractFilesDestinationDirectory(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf(" Problem in finding current working directory")
	}
	absPath, _ := filepath.Abs(pwd + "/../test/test-simple.zip")
	destinationPath, _ := filepath.Abs(pwd + "/../test/test.txt")
	expectedFile, _ := filepath.Abs(pwd + "/../test/test-simple/test.txt")

	Unzip(absPath, destinationPath)

	stat, err := os.Stat(expectedFile)
	if err != nil {
		t.Errorf("expected '%v' but got '%v'", true, err)
	}

	if stat.IsDir() {
		t.Errorf("expected '%v' but got '%v'", true, stat.IsDir())
	}

	if stat.Size() != 1 {
		t.Errorf("expected '%v' but got '%v'", 1, stat.Size())
	}

}

func TestFileNameWithoutExtension(t *testing.T) {
	name := FileNameWithoutExtension("abc.txt")

	assertEquals(name, "abc", t)
}
