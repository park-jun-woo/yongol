//ff:func feature=cli-init type=test-helper control=iteration dimension=1
//ff:what assertSkeletonFiles — verifies each expected skeleton file exists and is non-empty

package cliinit

import "testing"

func assertSkeletonFiles(t *testing.T, target string) {
	t.Helper()
	expectFiles := []string{
		"specs/manifest.yaml",
		"specs/api/openapi.yaml",
		"specs/db/sqlc.yaml",
		"specs/policy/authz.rego",
		"README.md",
		".gitignore",
	}
	for _, rel := range expectFiles {
		assertSkeletonFile(t, target, rel)
	}
}
