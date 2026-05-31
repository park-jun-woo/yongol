//ff:func feature=orchestrator type=test control=sequence
//ff:what TestDetectDirSSOT/directorySSOTs — glob 매칭 presence 결정 및 후보 목록 검증
package yongol

import (
	"testing"
)

func TestDetectDirSSOTGlobError(t *testing.T) {
	// A malformed glob pattern ("[") triggers filepath.ErrBadPattern,
	// exercising the hard-error branch.
	d := dirSSOT{kind: KindDDL, dir: t.TempDir(), globs: []string{"["}}
	_, err := detectDirSSOT(d)
	if err == nil {
		t.Fatal("expected error for malformed glob pattern")
	}
}
