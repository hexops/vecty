package vecty

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	// Try to remove all testdata/*.got.txt files now.
	matches, _ := filepath.Glob("testdata/*.got.txt")
	for _, match := range matches {
		os.Remove(match)
	}

	os.Exit(m.Run())
}
