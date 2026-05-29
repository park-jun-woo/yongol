//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-ddl
//ff:what canonicalDDLTableSet — ddlTableSet 키를 canonicalTableKey 로 정규화한 새 맵 반환

package ssac_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// canonicalDDLTableSet returns a new set whose keys are the DDL table names
// normalised through canonicalTableKey (singular lower-snake), so model↔table
// matching agrees with XSD-55. It never mutates the map returned by ddlTableSet
// (which may be the shared Ground lookup g.Lookup["DDL.table"]); a fresh map is
// always allocated.
func canonicalDDLTableSet(fs *yongol.Fullstack) map[string]bool {
	raw := ddlTableSet(fs)
	set := make(map[string]bool, len(raw))
	for name := range raw {
		set[canonicalTableKey(name)] = true
	}
	return set
}
