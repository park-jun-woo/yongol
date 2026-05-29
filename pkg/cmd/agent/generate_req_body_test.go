//ff:func feature=agent type=test control=sequence
//ff:what TestGenerateReqBody — body 불필요 method는 nil 반환, 필요 method는 LLM 에러 전파 검증

package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestGenerateReqBody(t *testing.T) {
	cfg := Config{Backend: "unsupported-backend", Model: "m"}

	// A GET-style op ("List...") needs no requestBody → returns nil, nil
	// without ever invoking the LLM.
	body, err := generateReqBody(features.Feature{Op: "ListUsers", Path: "/users"}, "", cfg)
	if err != nil || body != nil {
		t.Fatalf("ListUsers → %v, %v; want nil, nil", body, err)
	}

	// A POST-style op ("Create...") needs a body; with an unsupported backend
	// the LLM call fails and the error is wrapped/propagated.
	if _, err := generateReqBody(features.Feature{Op: "CreateUser", Path: "/users"}, "", cfg); err == nil {
		t.Error("expected error from LLM call for CreateUser")
	}
}
