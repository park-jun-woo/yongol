//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — design 토큰 참조/미지 prop 검사 헬퍼 직접 호출
package design

import (
	"testing"
)

func TestByNameResolveToken_ZeroCov(t *testing.T) {
	fs := designFS()
	if !resolveToken(fs, "colors.primary") {
		t.Errorf("colors.primary should resolve")
	}
	if !resolveToken(fs, "rounded.md") {
		t.Errorf("rounded.md should resolve")
	}
	if !resolveToken(fs, "spacing.sm") {
		t.Errorf("spacing.sm should resolve")
	}
	if resolveToken(fs, "colors.missing") {
		t.Errorf("missing color should not resolve")
	}
	if resolveToken(fs, "typography.body") {
		t.Errorf("typography branch nil map should not resolve")
	}
	if resolveToken(fs, "noseparator") {
		t.Errorf("no-dot ref should not resolve")
	}
	if resolveToken(fs, "unknown.group") {
		t.Errorf("unknown group should not resolve")
	}
}
