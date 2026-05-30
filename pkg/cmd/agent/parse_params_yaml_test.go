//ff:func feature=agent type=test control=sequence
//ff:what TestParseParamsYAML — 직접 배열/wrapping 키 해제/에러 케이스 파싱 검증

package agent

import "testing"

func TestParseParamsYAML(t *testing.T) {
	// Direct array form.
	arr, err := parseParamsYAML("- name: id\n  in: path\n- name: q\n  in: query")
	if err != nil {
		t.Fatalf("direct array: %v", err)
	}
	if len(arr) != 2 {
		t.Errorf("direct array len = %d, want 2", len(arr))
	}

	// Wrapped under 'parameters' key.
	arr, err = parseParamsYAML("parameters:\n  - name: id\n    in: path")
	if err != nil {
		t.Fatalf("wrapped: %v", err)
	}
	if len(arr) != 1 {
		t.Errorf("wrapped len = %d, want 1", len(arr))
	}

	// Map without 'parameters' key is an error.
	if _, err := parseParamsYAML("other:\n  - a"); err == nil {
		t.Error("expected error for map missing 'parameters' key")
	}

	// 'parameters' that is not an array is an error.
	if _, err := parseParamsYAML("parameters: scalar"); err == nil {
		t.Error("expected error when 'parameters' is not an array")
	}

	// Completely malformed YAML fails both the array and the map unmarshal.
	if _, err := parseParamsYAML("a: [1, 2"); err == nil {
		t.Error("expected error for malformed YAML")
	}
}
