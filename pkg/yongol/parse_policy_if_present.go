//ff:func feature=orchestrator type=loader control=sequence
//ff:what Rego 정책 탐지 시 @ownership/allow 메타데이터 구조화 파싱
package yongol

import (
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

// parsePolicyIfPresent parses Rego policies when KindPolicy is present.
// ParsePolicies uses regex-based parsing over Rego to structure SSOT
// metadata such as @ownership and allow rules. The results are consumed by
// Ground/validator downstream.
// (The former rego.ParseDir OPA AST call has been removed because its
// results were not consumed anywhere.)
func parsePolicyIfPresent(fs *Fullstack, has map[SSOTKind]DetectedSSOT) {
	d, ok := has[KindPolicy]
	if !ok {
		return
	}
	policies, pdiags := rego.ParsePolicies(d.Path)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, pdiags...)
	if len(pdiags) == 0 {
		fs.ParsedPolicies = policies
	}
}
