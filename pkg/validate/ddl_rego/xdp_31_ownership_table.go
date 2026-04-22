//ff:func feature=validate type=rule control=iteration dimension=2 topic=policy-check
//ff:what XDP-31 — @ownership table → DDL

package ddl_rego

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdp31OwnershipTable validates XDP-31: Rego @ownership 어노테이션이 참조하는
// 테이블이 DDL 에 존재하는지 확인한다.
func xdp31OwnershipTable(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	tables := g.Lookup["DDL.table"]
	if tables == nil {
		// Ground 미생성은 상위 파이프라인 버그. XDP-31 은 Ground 단일 소스만 사용한다.
		return nil
	}

	seen := make(map[string]bool)
	var diags []diagnostic.Diagnostic
	for _, p := range fs.ParsedPolicies {
		for _, om := range p.Ownerships {
			if om.Table == "" {
				continue
			}
			key := p.File + "|" + om.Resource + "|" + om.Table
			if seen[key] {
				continue
			}
			seen[key] = true
			if !tables[om.Table] {
				diags = append(diags, diagnostic.Diagnostic{
					File:  p.File,
					Line:  om.SourceLine,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf(
						"[XDP-31] @ownership %s 의 table %q 가 DDL 에 존재하지 않습니다",
						om.Resource, om.Table),
					Advice: fmt.Sprintf("DDL 에 테이블 %s 를 정의하거나 Rego @ownership 에서 제거하세요", om.Table),
				})
			}
		}
	}
	return diags
}
