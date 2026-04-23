//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseDir — 존재하지 않는 디렉토리는 빈 결과 반환

package sqlc

import "testing"

func TestParseDir_MissingDir(t *testing.T) {
	// non-existent directory is not an error (empty result)
	specs, diags := ParseDir("/nonexistent/dir/for/sqlc/parse/dir")
	if len(specs) != 0 {
		t.Errorf("want 0 specs, got %d", len(specs))
	}
	if len(diags) != 0 {
		t.Errorf("want 0 diags, got %d: %v", len(diags), diags)
	}
}
