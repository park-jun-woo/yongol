//ff:func feature=cli-init type=test-helper control=sequence
//ff:what assertSkeletonOpenAPI — openapi.yaml contains empty paths placeholder

package cliinit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func assertSkeletonOpenAPI(t *testing.T, target string) {
	t.Helper()
	openapi, err := os.ReadFile(filepath.Join(target, "specs/api/openapi.yaml"))
	if err != nil {
		t.Fatalf("read openapi: %v", err)
	}
	if !strings.Contains(string(openapi), "paths: {}") {
		t.Errorf("openapi missing empty paths placeholder")
	}
}
