package unzip

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

	t.Cleanup(func() {
		if err == nil {
			os.RemoveAll(destinationPath)
		}
	})

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

	t.Cleanup(func() {
		if err == nil {
			os.RemoveAll(destinationPath)
		}
	})

}

func TestUnzipExtractFilesDestinationDirectory(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf(" Problem in finding current working directory")
	}
	absPath, _ := filepath.Abs(pwd + "/../test/test-simple.zip")
	destinationPath, _ := filepath.Abs(pwd + "/../test/test-simple")
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

	t.Cleanup(func() {
		if err == nil {
			os.RemoveAll(destinationPath)
		}
	})

}

func TestUnzipExtractNestedFilesDestinationDirectory(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf(" Problem in finding current working directory")
	}
	absPath, _ := filepath.Abs(pwd + "/../test/test-nested.zip")
	destinationPath, _ := filepath.Abs(pwd + "/../test/test-nested")
	expectedFile, _ := filepath.Abs(pwd + "/../test/test-nested/dir/test.txt")

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

	t.Cleanup(func() {
		if err == nil {
			os.RemoveAll(destinationPath)
		}
	})
}

func TestFileNameWithoutExtension(t *testing.T) {
	name := FileNameWithoutExtension("abc.txt")

	assertEquals(name, "abc", t)
}
