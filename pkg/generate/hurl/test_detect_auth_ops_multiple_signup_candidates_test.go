//ff:func feature=gen-hurl type=test control=iteration dimension=1
//ff:what TestDetectAuthOpsMultipleSignupCandidates — 다수 signup 후보 → 알파벳 우선

package hurl

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestDetectAuthOpsMultipleSignupCandidates pins the deterministic
// picker: two signup-shape ops → alphabetical-first wins + WARNING.
func TestDetectAuthOpsMultipleSignupCandidates(t *testing.T) {
	admin := newPasswordOp("AdminSignup", true, true)
	user := newPasswordOp("UserSignup", true, true)
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/auth/admin-signup", &openapi3.PathItem{Post: admin}),
		openapi3.WithPath("/auth/user-signup", &openapi3.PathItem{Post: user}),
	)}
	adminFn := signupServiceFunc("AdminSignup")
	userFn := signupServiceFunc("UserSignup")
	fs := &yongol.Fullstack{OpenAPIDoc: doc, ServiceFuncs: []ssac.ServiceFunc{adminFn, userFn}}
	signup, _, warns := detectAuthOps(fs)
	if signup == nil {
		t.Fatalf("expected a signup pick, got nil (warns=%v)", warns)
	}
	if signup.OpID != "AdminSignup" {
		t.Errorf("signup pick: want AdminSignup (alphabetical first), got %q", signup.OpID)
	}
	sawMulti := false
	for _, w := range warns {
		if strings.Contains(w, "multiple signup candidates") {
			sawMulti = true
			break
		}
	}
	if !sawMulti {
		t.Errorf("expected 'multiple signup candidates' warning; got %v", warns)
	}
}
