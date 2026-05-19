//ff:func feature=cli-init type=test control=sequence
//ff:what TestRunCreatesSkeleton — Run writes the expected skeleton tree

package cliinit

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestRunCreatesSkeleton(t *testing.T) {
	featPath := writeTempFeatures(t)
	tmp := t.TempDir()
	target := filepath.Join(tmp, "myapp")
	var outBuf, errBuf bytes.Buffer
	err := Run(&outBuf, &errBuf, Options{
		ProjectID:    "Myapp",
		FeaturesPath: featPath,
		Description:  "Test description",
		Dir:          target,
		Module:       "github.com/test/myapp",
	})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	assertSkeletonFiles(t, target)
	assertSkeletonDirs(t, target)
	assertSkeletonManifest(t, target)
	assertSkeletonSqlc(t, target)
}
