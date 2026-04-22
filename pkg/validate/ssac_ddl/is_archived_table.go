//ff:func feature=validate type=util control=sequence topic=ssac-ddl
//ff:what isArchivedTable — Ground.Flags["archived.<table>"] 조회

package ssac_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// isArchivedTable reports whether the DDL table is flagged @archived.
func isArchivedTable(fs *yongol.Fullstack, table string) bool {
	g := fs.Ground()
	if g == nil {
		return false
	}
	return g.Flags["archived."+table]
}
