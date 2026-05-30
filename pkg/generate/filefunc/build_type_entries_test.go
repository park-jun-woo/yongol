//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestBuildTypeEntries — 고정 type 카테고리 맵 반환 검증

package filefunc

import "testing"

func TestBuildTypeEntries(t *testing.T) {
	got := buildTypeEntries()
	want := []string{
		"handler", "service", "model", "query", "middleware", "config",
		"accessor", "util", "generator", "loader", "command", "test",
		"test-helper",
	}
	if len(got) != len(want) {
		t.Errorf("expected %d types, got %d: %v", len(want), len(got), got)
	}
	for _, k := range want {
		if desc, ok := got[k]; !ok || desc == "" {
			t.Errorf("missing or empty type %q: %q", k, desc)
		}
	}
}
