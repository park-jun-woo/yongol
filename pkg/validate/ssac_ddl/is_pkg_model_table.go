//ff:func feature=validate type=util control=sequence topic=ssac-ddl
//ff:what isPkgModelTable — Ground.Flags["pkgModel.<table>"] 조회 (미적재 시 false)

package ssac_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// isPkgModelTable reports whether the DDL table is owned by a pkg/* built-in model
// (e.g. pkg/session, pkg/cache, pkg/file). Lookup key follows the per-target
// "pkgModel.<table>" convention so future populators can opt in per table.
func isPkgModelTable(fs *yongol.Fullstack, table string) bool {
	g := fs.Ground()
	if g == nil {
		return false
	}
	return g.Flags["pkgModel."+table]
}
