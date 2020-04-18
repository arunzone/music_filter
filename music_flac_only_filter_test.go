package main

import (
	"testing"
)

func TestIsHiddenFile(t *testing.T) {
	isHiddenFile := IsHiddenFile(".DS_Store")

	if !isHiddenFile {
		t.Errorf("expected '%v' but got '%v'", true, isHiddenFile)
	}
}