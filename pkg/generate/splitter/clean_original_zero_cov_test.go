//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what zz_zerocov_test — splitter 패키지의 0% 커버리지 함수(cleanOriginal/preserveComments/isPreservedFile/SplitDirectory) 단위 테스트
package splitter

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCleanOriginal_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	// sqlc tool: models.go and foo.sql.go are originals; querier.go preserved;
	// split.go in keep; nested dir is skipped.
	writeFileZeroCov(t, dir, "models.go", "package p\n")
	writeFileZeroCov(t, dir, "foo.sql.go", "package p\n")
	writeFileZeroCov(t, dir, "querier.go", "package p\n")
	writeFileZeroCov(t, dir, "user.model.go", "package p\n") // a split result, in keep
	writeFileZeroCov(t, dir, "unrelated.txt", "x\n")
	if err := os.Mkdir(filepath.Join(dir, "nested"), 0o755); err != nil {
		t.Fatal(err)
	}

	keep := map[string]bool{"user.model.go": true}
	if err := cleanOriginal(dir, ToolSQLC, keep); err != nil {
		t.Fatalf("cleanOriginal: %v", err)
	}

	mustGone := []string{"models.go", "foo.sql.go"}
	for _, n := range mustGone {
		if _, err := os.Stat(filepath.Join(dir, n)); !os.IsNotExist(err) {
			t.Errorf("%s should have been removed", n)
		}
	}
	mustStay := []string{"querier.go", "user.model.go", "unrelated.txt"}
	for _, n := range mustStay {
		if _, err := os.Stat(filepath.Join(dir, n)); err != nil {
			t.Errorf("%s should have survived: %v", n, err)
		}
	}
}
