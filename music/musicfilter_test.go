package music

import (
	"os"
	"testing"
)

func TestAbsolutePath(t *testing.T) {
	absolutePath := AbsolutePath("a", "b")
	expected := "a/b"

	if absolutePath != expected {
		t.Errorf("expected '%v' but got '%v'", expected, absolutePath)
	}
}

func TestUnzipNotReturnError(t *testing.T) {
	getwd, err := os.Getwd()
	if err != nil {
		t.Fatalf(" Problem in finding current working directory")
	}
	status := Unzip(getwd + "/../test/test.txt.zip")

	if status != nil {
		t.Errorf("expected '%v' but got '%v'", nil, status)
	}
}
