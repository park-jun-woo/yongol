//ff:func feature=features-parse type=test control=sequence
//ff:what TestByName_ZeroCov — features 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속

package features

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestByNameCollectSequenceLines_ZeroCov(t *testing.T) {
	var root yaml.Node
	if err := yaml.Unmarshal([]byte("- a\n- b\n"), &root); err != nil {
		t.Fatal(err)
	}
	seq := root.Content[0]
	lines := collectSequenceLines(seq)
	if len(lines) != 2 {
		t.Errorf("collectSequenceLines = %v, want 2", lines)
	}
}

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
