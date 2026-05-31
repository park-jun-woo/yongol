//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what matchOwnershipIR -- OwnershipMapping 목록에서 resource 매칭 OwnershipInfo 구성

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// matchOwnershipIR returns OwnershipInfo for the first ownership mapping whose
// Resource equals resource, resolving the resource PK from DDL when available.
// Returns nil when no mapping matches.
func matchOwnershipIR(fs *yongol.Fullstack, mappings []rego.OwnershipMapping, resource string) *OwnershipInfo {
	for _, om := range mappings {
		if om.Resource != resource {
			continue
		}
		info := OwnershipInfo{Table: om.Table, OwnerColumn: om.Column}
		if len(fs.DDLTables) > 0 {
			info.ResourcePK = findTablePK(fs, om.Table)
		}
		return &info
	}
	return nil
}
