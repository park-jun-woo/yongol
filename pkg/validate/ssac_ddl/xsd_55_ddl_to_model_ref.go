//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-ddl
//ff:what XSD-55 — DDL table → SSaC modelRef

package ssac_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsd55DDLToModelRef validates XSD-55: every DDL table must be referenced by
// some SSaC @model or @result. Archived and pkgModel tables are exempt.
func xsd55DDLToModelRef(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	referenced := buildReferencedTables(fs.ServiceFuncs)
	var diags []diagnostic.Diagnostic
	for _, t := range fs.DDLTables {
		if referenced[canonicalTableKey(t.Name)] {
			continue
		}
		if isArchivedTable(fs, t.Name) {
			continue
		}
		if isFuncManagedTable(fs, t.Name) {
			continue
		}
		if isPkgModelTable(fs, t.Name) {
			continue
		}
		if isAuthRequiredTable(fs, t.Name) {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    t.File,
			Line:    t.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[XSD-55] DDL table %q is not referenced by any SSaC @model or @result", t.Name),
			Advice:  fmt.Sprintf("DDL 테이블 %s 를 어떤 SSaC @model 또는 @result 에서 사용하거나 DDL 에서 제거하세요. RPC/함수가 관리하는 활성 테이블이면 -- @func-managed, 미사용/폐기 테이블이면 -- @archived 어노테이션을 CREATE TABLE 바로 위에 추가하세요", t.Name),
		})
	}
	return diags
}
