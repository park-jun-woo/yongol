//ff:func feature=agent type=test control=sequence
//ff:what TestParseReqBodyYAML — requestBody wrapping 키 해제, 직접 맵, 에러 케이스 검증

package agent

import "testing"

func TestParseReqBodyYAML(t *testing.T) {
	// Wrapped single 'requestBody' key is unwrapped.
	m, err := parseReqBodyYAML("requestBody:\n  required: true")
	if err != nil {
		t.Fatalf("wrapped: %v", err)
	}
	if m["required"] != true {
		t.Errorf("unwrapped map = %v, want required:true", m)
	}

	// Direct map (no wrapping) is returned as-is.
	m, err = parseReqBodyYAML("content:\n  application/json: {}")
	if err != nil {
		t.Fatalf("direct: %v", err)
	}
	if _, ok := m["content"]; !ok {
		t.Errorf("direct map = %v, want content key", m)
	}

	// Non-map YAML is an error.
	if _, err := parseReqBodyYAML("- not a map"); err == nil {
		t.Error("expected error for non-map YAML")
	}
}
