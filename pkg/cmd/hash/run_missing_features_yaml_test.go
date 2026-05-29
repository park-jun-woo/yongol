//ff:func feature=cli-hash type=test control=sequence
//ff:what TestRun_MissingFeaturesYaml — features.yaml 없을 때 에러 반환 확인

package clihash

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun_MissingFeaturesYaml(t *testing.T) {
	dir := t.TempDir()

	var buf bytes.Buffer
	err := Run(&buf, dir)
	if err == nil {
		t.Fatal("expected error for missing features.yaml, got nil")
	}
	if !strings.Contains(err.Error(), "features.yaml") {
		t.Errorf("error should mention features.yaml: %v", err)
	}
}
