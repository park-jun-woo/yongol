//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what isPgtypeOverrideFor — sqlc override entry 가 (dbType, pgPackage, pgType) 매핑인지 판정

package query

import "strings"

// isPgtypeOverrideFor reports whether ov targets PG dbType and maps it
// to the canonical pgx/v5 pgtype Go type (pgPackage.pgType from the
// import "github.com/jackc/pgx/v5/pgtype").
func isPgtypeOverrideFor(ov sqlcOverrideEntry, dbType, pgPackage, pgType string) bool {
	if !strings.EqualFold(ov.DBType, dbType) {
		return false
	}
	if ov.GoType.Package != pgPackage || ov.GoType.Type != pgType {
		return false
	}
	return ov.GoType.Import == "github.com/jackc/pgx/v5/pgtype"
}
