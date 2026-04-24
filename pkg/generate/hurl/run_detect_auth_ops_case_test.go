//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what runDetectAuthOpsCase — TestDetectAuthOps 단일 케이스 실행

package hurl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runDetectAuthOpsCase executes one detectAuthOps fixture and asserts the
// resulting role + expected warning substring. Extracted from the outer
// test so the range body stays within the Q4 PURE budget.
func runDetectAuthOpsCase(t *testing.T, tc detectAuthOpsFixture) {
	t.Helper()
	op := newPasswordOp(tc.opID, tc.public, tc.hasPassword)
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/auth/"+tc.opID, &openapi3.PathItem{Post: op}),
	)}
	var funcs []ssac.ServiceFunc
	if fn := tc.funcBuilder(tc.opID); fn != nil {
		funcs = []ssac.ServiceFunc{*fn}
	}
	fs := &yongol.Fullstack{OpenAPIDoc: doc, ServiceFuncs: funcs}
	signup, login, warns := detectAuthOps(fs)
	assertDetectAuthOpsRole(t, tc, signup, login, warns)
	if tc.wantWarnSub != "" {
		assertWarningContains(t, warns, tc.wantWarnSub)
	}
}
