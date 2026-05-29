//ff:func feature=agent type=test control=sequence
//ff:what TestParseSchema200YAML — schema wrapping 키 해제, 직접 맵, 에러 케이스 검증

package agent

import "testing"

func TestParseSchema200YAML(t *testing.T) {
	// Wrapped single 'schema' key is unwrapped.
	m, err := parseSchema200YAML("schema:\n  type: object")
	if err != nil {
		t.Fatalf("wrapped: %v", err)
	}
	if m["type"] != "object" {
		t.Errorf("unwrapped = %v, want type:object", m)
	}

	// Direct map is returned as-is.
	m, err = parseSchema200YAML("type: string")
	if err != nil {
		t.Fatalf("direct: %v", err)
	}
	if m["type"] != "string" {
		t.Errorf("direct = %v, want type:string", m)
	}

	// Non-map YAML is an error.
	if _, err := parseSchema200YAML("- a\n- b"); err == nil {
		t.Error("expected error for non-map YAML")
	}
}
