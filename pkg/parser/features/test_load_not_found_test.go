//ff:func feature=features type=test control=sequence
//ff:what Load — features.yaml 없을 때 에러 진단 테스트
package features

import (
	"testing"
)

func TestLoad_NotFound(t *testing.T) {
	dir := t.TempDir()
	_, diags := Load(dir)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if diags[0].Level != "ERROR" {
		t.Errorf("want ERROR, got %s", diags[0].Level)
	}
}
