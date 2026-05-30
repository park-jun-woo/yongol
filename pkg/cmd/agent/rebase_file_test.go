//ff:func feature=agent type=test control=sequence
//ff:what TestRebaseFile — 절대 경로를 specsDir 상대 경로로 변환, 상대 경로는 그대로 반환 검증

package agent

import "testing"

func TestRebaseFile(t *testing.T) {
	// Relative path is returned unchanged.
	if got := rebaseFile("db/schema.sql", "/specs"); got != "db/schema.sql" {
		t.Errorf("relative = %q, want db/schema.sql", got)
	}
	// Absolute path under specs is rebased and slash-normalized.
	if got := rebaseFile("/specs/api/openapi.yaml", "/specs"); got != "api/openapi.yaml" {
		t.Errorf("absolute = %q, want api/openapi.yaml", got)
	}
	// An absolute file with a *relative* base makes filepath.Rel fail, so the
	// original (absolute) path is returned unchanged.
	abs := "/specs/api/openapi.yaml"
	if got := rebaseFile(abs, "relative/base"); got != abs {
		t.Errorf("rel error fallback = %q, want %q", got, abs)
	}
}
