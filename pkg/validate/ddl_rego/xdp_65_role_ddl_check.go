//ff:func feature=validate type=rule control=iteration dimension=2 topic=policy-check
//ff:what XDP-65 — Rego role → DDL CHECK

package ddl_rego

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdp65RoleDDLCheck validates XDP-65: Rego 정책이 비교하는 role 값들이
// DDL 의 role 컬럼 CHECK(role IN (...)) 제약에 모두 정의되어 있는지
// 확인한다. DDL 에 role CHECK 가 전혀 없으면 사용자 role 모델이 없는
// 것으로 간주하고 통과시킨다 (bak/check_rego_role_ddl.go 와 동일 정책).
func xdp65RoleDDLCheck(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}

	// Collect allowed role values from any DDL table's CHECK on "role" column.
	allowed := make(map[string]bool)
	for _, t := range fs.DDLTables {
		if vals, ok := t.CheckEnums["role"]; ok {
			for _, v := range vals {
				allowed[v] = true
			}
		}
	}
	if len(allowed) == 0 {
		return nil
	}

	// Collect Rego role values with source context (first occurrence wins).
	type ctx struct {
		file string
		line int
	}
	regoRoles := make(map[string]ctx)
	for _, p := range fs.ParsedPolicies {
		for _, r := range p.Rules {
			if !r.UsesRole || r.RoleValue == "" {
				continue
			}
			if _, exists := regoRoles[r.RoleValue]; exists {
				continue
			}
			regoRoles[r.RoleValue] = ctx{file: p.File, line: r.SourceLine}
		}
	}

	var diags []diagnostic.Diagnostic
	for rv, c := range regoRoles {
		if allowed[rv] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:  c.file,
			Line:  c.line,
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelError,
			Message: fmt.Sprintf(
				"[XDP-65] Rego role %q 가 DDL CHECK 제약에 정의되어 있지 않습니다",
				rv),
			Advice: fmt.Sprintf("DDL 사용자 테이블의 role 컬럼 CHECK IN 에 '%s' 를 추가하세요", rv),
		})
	}
	return diags
}
