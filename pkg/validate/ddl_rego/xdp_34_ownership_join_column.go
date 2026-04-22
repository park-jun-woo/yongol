//ff:func feature=validate type=rule control=iteration dimension=2 topic=policy-check
//ff:what XDP-34 — @ownership via join column → DDL

package ddl_rego

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdp34OwnershipJoinColumn validates XDP-34: Rego @ownership via 의
// 조인 FK 컬럼이 조인 테이블의 DDL 컬럼 정의에 존재하는지 확인한다.
// join table 이 DDL 에 없으면 XDP-33 에서 보고되므로 여기서는 건너뛴다.
func xdp34OwnershipJoinColumn(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	g := fs.Ground()
	tables := buildDDLTableSet(fs, g)
	columnsByTable := buildDDLColumnIndex(fs)

	seen := make(map[string]bool)
	var diags []diagnostic.Diagnostic
	for _, p := range fs.ParsedPolicies {
		for _, om := range p.Ownerships {
			if om.JoinTable == "" || om.JoinFK == "" {
				continue
			}
			if !tables[om.JoinTable] {
				continue // reported by XDP-33
			}
			key := p.File + "|" + om.Resource + "|" + om.JoinTable + "." + om.JoinFK
			if seen[key] {
				continue
			}
			seen[key] = true
			cols := columnsByTable[om.JoinTable]
			if !cols[om.JoinFK] {
				diags = append(diags, diagnostic.Diagnostic{
					File:  p.File,
					Line:  om.SourceLine,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf(
						"[XDP-34] @ownership %s via 의 join 컬럼 %s.%s 가 DDL 에 존재하지 않습니다",
						om.Resource, om.JoinTable, om.JoinFK),
					Advice: fmt.Sprintf("DDL join 테이블 %s 에 컬럼 %s 를 추가하세요", om.JoinTable, om.JoinFK),
				})
			}
		}
	}
	return diags
}
