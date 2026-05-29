//ff:func feature=cli type=test-helper control=sequence
//ff:what test: mustWriteEmpty — helper that creates an empty file for testing

package main

import (
	"os"
	"testing"
)

// mustWriteEmpty creates an empty file at path and fatally fails the
// test on error. Extracted from test files so multiple _test.go cases
// can share a single fixture helper without duplicating boilerplate.
func mustWriteEmpty(t *testing.T, path string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(""), 0o644); err != nil {
		t.Fatal(err)
	}
}
