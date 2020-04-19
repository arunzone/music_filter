package file

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopy(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf(" Problem in finding current working directory")
	}
	copyPath, _ := filepath.Abs(pwd + "/../test/copy")
	testcsv := filepath.Join(copyPath, "test.csv")
	music := filepath.Join(copyPath, "Music")


	Copy([]string{testcsv}, music)

	stat, err := os.Stat(filepath.Join(music, "test.csv"))
	if err != nil || stat.Size() != 5 {
		t.Errorf("expected '%v' but got '%v'", stat.Size(), 5)
	}

	t.Cleanup(func() {
		if err == nil {
			os.RemoveAll(music)
		}
	})
}