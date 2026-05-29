//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestMirrorSpecsTests_EmptySpecsDir — specsDir="" 시 no-op

package hurl_mirror

import "testing"

// TestMirrorSpecsTests_EmptySpecsDir verifies that specsDir="" short
// circuits without touching the filesystem.
func TestMirrorSpecsTests_EmptySpecsDir(t *testing.T) {
	n, err := MirrorSpecsTests("", "")
	if err != nil {
		t.Fatalf("MirrorSpecsTests: %v", err)
	}
	if n != 0 {
		t.Fatalf("mirrored = %d; want 0", n)
	}
}
