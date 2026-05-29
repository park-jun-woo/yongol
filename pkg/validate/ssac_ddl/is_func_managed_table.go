//ff:func feature=validate type=util control=sequence topic=ssac-ddl
//ff:what isFuncManagedTable — Ground.Flags["funcManaged.<table>"] 조회

package ssac_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// isFuncManagedTable reports whether the DDL table is flagged `-- @func-managed`,
// i.e. actively managed by a @call'd function/RPC. XSD-55 exempts such tables.
func isFuncManagedTable(fs *yongol.Fullstack, table string) bool {
	g := fs.Ground()
	if g == nil {
		return false
	}
	return g.Flags["funcManaged."+table]
}
