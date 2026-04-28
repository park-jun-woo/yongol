//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-infra
//ff:what collectDDLNames — fs.DDLTables 의 name 세트 구성

package manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

// collectDDLNames returns the set of DDL table names present in fs so
// validateBuiltinBackend can check required tables in O(1).
func collectDDLNames(fs *yongol.Fullstack) map[string]bool {
	out := make(map[string]bool, len(fs.DDLTables))
	for _, t := range fs.DDLTables {
		out[t.Name] = true
	}
	return out
}
