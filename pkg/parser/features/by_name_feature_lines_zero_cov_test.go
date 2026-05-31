//ff:func feature=features type=test control=sequence
//ff:what TestByName_ZeroCov — features 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package features

import (
	"testing"
)

func TestByNameFeatureLines_ZeroCov(t *testing.T) {
	data := []byte("features:\n  - one\n  - two\n  - three\n")
	lines := extractFeatureLines(data)
	if len(lines) != 3 {
		t.Errorf("extractFeatureLines = %v, want 3", lines)
	}

	// no features key.
	if got := extractFeatureLines([]byte("other: x\n")); got != nil {
		t.Errorf("extractFeatureLines without features should be nil, got %v", got)
	}

	// invalid yaml.
	if got := extractFeatureLines([]byte("\t- bad\n")); got != nil {
		t.Errorf("extractFeatureLines invalid yaml should be nil, got %v", got)
	}
}
