//ff:func feature=policy type=parser control=iteration dimension=1
//ff:what ParsePolicies — 디렉토리 내 .rego 파일에서 Policy 목록 추출
package rego

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParsePolicies parses all .rego files in dir and returns structured policies.
func ParsePolicies(dir string) ([]Policy, []diagnostic.Diagnostic) {
	matches, err := filepath.Glob(filepath.Join(dir, "*.rego"))
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    dir,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "[Rego glob] cannot glob rego dir: " + err.Error(),
		}}
	}
	var policies []Policy
	for _, path := range matches {
		p, diags := ParsePolicyFile(path)
		if len(diags) > 0 {
			return nil, diags
		}
		policies = append(policies, *p)
	}
	return policies, nil
}
