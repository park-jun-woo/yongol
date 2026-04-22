//ff:func feature=validate type=rule control=iteration dimension=2 topic=policy-check
//ff:what XDP-32 — @ownership column → DDL

package ddl_rego

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdp32OwnershipColumn validates XDP-32: Rego @ownership 이 참조하는
// 컬럼이 대상 테이블의 DDL 컬럼 정의에 존재하는지 확인한다. 테이블이
// 존재하지 않으면 XDP-31 에서 보고하므로 여기서는 건너뛴다.
func xdp32OwnershipColumn(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	g := fs.Ground()

	columnsByTable := buildDDLColumnIndex(fs)
	tables := buildDDLTableSet(fs, g)

	seen := make(map[string]bool)
	var diags []diagnostic.Diagnostic
	for _, p := range fs.ParsedPolicies {
		for _, om := range p.Ownerships {
			if om.Table == "" || om.Column == "" {
				continue
			}
			if !tables[om.Table] {
				continue // reported by XDP-31
			}
			key := p.File + "|" + om.Resource + "|" + om.Table + "." + om.Column
			if seen[key] {
				continue
			}
			seen[key] = true
			cols := columnsByTable[om.Table]
			if !cols[om.Column] {
				diags = append(diags, diagnostic.Diagnostic{
					File:  p.File,
					Line:  om.SourceLine,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf(
						"[XDP-32] @ownership %s 의 컬럼 %s.%s 가 DDL 에 존재하지 않습니다",
						om.Resource, om.Table, om.Column),
					Advice: fmt.Sprintf("DDL 테이블 %s 에 컬럼 %s 를 추가하세요 (보통 owner_id, user_id 등)", om.Table, om.Column),
				})
			}
		}
	}
	return diags
}
