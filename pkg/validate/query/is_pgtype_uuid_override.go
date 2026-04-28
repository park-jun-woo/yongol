//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what isPgtypeUUIDOverride — sqlc override entry 가 PG uuid → pgtype.UUID 매핑인지 판정

package query

import "strings"

// isPgtypeUUIDOverride reports whether the override targets PG `uuid` and
// maps it to `github.com/jackc/pgx/v5/pgtype.UUID`. Used by Q-12 when
// counting how many of the two required (NULL / NOT NULL) entries are
// present.
func isPgtypeUUIDOverride(ov sqlcOverrideEntry) bool {
	if !strings.EqualFold(ov.DBType, "uuid") {
		return false
	}
	if ov.GoType.Package != "pgtype" || ov.GoType.Type != "UUID" {
		return false
	}
	return ov.GoType.Import == "github.com/jackc/pgx/v5/pgtype"
}
