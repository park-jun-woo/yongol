//ff:func feature=validate type=rule control=iteration dimension=1 topic=rego-structural
//ff:what XPP-30 — Rego 정책에서 resource_owner 참조 시 @ownership 어노테이션 필수

package rego

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xpp30OwnershipNoAnnotation flags policies that reference `resource_owner`
// without declaring a corresponding @ownership mapping.
func xpp30OwnershipNoAnnotation(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, p := range fs.ParsedPolicies {
		if !usesResourceOwner(p.Rules) {
			continue
		}
		if len(p.Ownerships) == 0 {
			diags = append(diags, diagnostic.Diagnostic{
				File:    p.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[XPP-30] policy references resource_owner but declares no @ownership",
				Advice:  "정책 패키지 상단에 # @ownership table.column [join_table.fk] 를 추가하세요",
			})
		}
	}
	return diags
}
