//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — normPath / findDesignFiles 직접 호출
package design_manifest

import (
	"testing"
)

func TestNormPath_ZeroCov(t *testing.T) {
	if got := normPath("a\\b\\c"); got != "a/b/c" {
		t.Errorf("normPath = %q", got)
	}
	if got := normPath("a/./b/"); got != "a/b" {
		t.Errorf("normPath clean = %q", got)
	}
}
