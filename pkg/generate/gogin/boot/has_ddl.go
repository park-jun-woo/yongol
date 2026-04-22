//ff:func feature=gen-gogin type=util control=sequence
//ff:what hasDDL — Fullstack 에 DDL 이 파싱되어 존재하는지 여부

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasDDL returns true when the project has any DDL declared (either parsed
// SQL AST results or DDLTables). Readiness DB ping hinges on this.
func hasDDL(fs *yongol.Fullstack) bool {
	if fs == nil {
		return false
	}
	return len(fs.DDLResults) > 0 || len(fs.DDLTables) > 0
}
