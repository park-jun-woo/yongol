//ff:func feature=validate type=rule control=iteration dimension=2 topic=policy-check
//ff:what XDP-33 — @ownership via join table → DDL

package ddl_rego

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdp33OwnershipJoinTable validates XDP-33: Rego @ownership via 로 지정한
// 조인 테이블이 DDL 에 존재하는지 확인한다.
func xdp33OwnershipJoinTable(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	g := fs.Ground()
	tables := buildDDLTableSet(fs, g)

	seen := make(map[string]bool)
	var diags []diagnostic.Diagnostic
	for _, p := range fs.ParsedPolicies {
		for _, om := range p.Ownerships {
			if om.JoinTable == "" {
				continue
			}
			key := p.File + "|" + om.Resource + "|" + om.JoinTable
			if seen[key] {
				continue
			}
			seen[key] = true
			if !tables[om.JoinTable] {
				diags = append(diags, diagnostic.Diagnostic{
					File:  p.File,
					Line:  om.SourceLine,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf(
						"[XDP-33] @ownership %s via 의 join table %q 가 DDL 에 존재하지 않습니다",
						om.Resource, om.JoinTable),
					Advice: fmt.Sprintf("DDL 에 join 테이블 %s 를 정의하세요", om.JoinTable),
				})
			}
		}
	}
	return diags
}
