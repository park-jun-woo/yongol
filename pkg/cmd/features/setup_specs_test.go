//ff:func feature=features type=test control=sequence
//ff:what TestRunRemove — 빈 ops/존재안함/abort/확인삭제/ssac-skip/--yes 성공 분기 검증
package features

import (
	"os"
	"path/filepath"
	"testing"
)

func setupSpecs(t *testing.T) string {
	t.Helper()
	specs := t.TempDir()
	if err := os.WriteFile(filepath.Join(specs, "features.yaml"), []byte(twoFeats), 0o644); err != nil {
		t.Fatal(err)
	}
	return specs
}
