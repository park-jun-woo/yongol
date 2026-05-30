//ff:func feature=cli-init type=test control=sequence
//ff:what TestSkeletonDirs — 필수 디렉토리 목록 포함 및 shallow→deep 순서 검증

package cliinit

import "testing"

func TestSkeletonDirs(t *testing.T) {
	dirs := skeletonDirs()
	if len(dirs) == 0 {
		t.Fatal("expected non-empty skeleton dir list")
	}
	want := map[string]bool{
		"specs":                      false,
		"specs/db/queries":           false,
		"specs/frontend/components":  false,
		"specs/tests":                false,
	}
	for _, d := range dirs {
		if _, ok := want[d]; ok {
			want[d] = true
		}
	}
	for d, seen := range want {
		if !seen {
			t.Errorf("skeletonDirs missing %q", d)
		}
	}
	// "specs" must precede its children (shallow->deep ordering).
	if dirs[0] != "specs" {
		t.Errorf("first dir = %q, want specs", dirs[0])
	}
}
