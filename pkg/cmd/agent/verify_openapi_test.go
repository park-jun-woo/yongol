//ff:func feature=agent type=test control=sequence
//ff:what TestVerifyOpenAPI — 유효한 OpenAPI 문서 통과, 잘못된 YAML/스펙 에러 검증

package agent

import "testing"

func TestVerifyOpenAPI(t *testing.T) {
	valid := []byte(`openapi: 3.0.0
info:
  title: t
  version: "1.0"
paths: {}
`)
	if err := verifyOpenAPI(valid); err != nil {
		t.Errorf("valid doc rejected: %v", err)
	}

	// Malformed YAML cannot be loaded.
	if err := verifyOpenAPI([]byte(":::not yaml:::")); err == nil {
		t.Error("expected error for malformed YAML")
	}

	// Structurally invalid OpenAPI (missing required info/version) fails validation.
	if err := verifyOpenAPI([]byte("openapi: 3.0.0\npaths: {}\n")); err == nil {
		t.Error("expected validation error for missing info block")
	}
}
