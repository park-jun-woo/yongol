//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what isPgtypeOverrideFor — sqlc override entry 가 (dbType, pgPackage, pgType) 매핑인지 판정

package query

import "strings"

// isPgtypeOverrideFor reports whether ov targets PG dbType and maps it
// to the canonical pgx/v5 pgtype Go type (pgPackage.pgType from the
// import "github.com/jackc/pgx/v5/pgtype"). The `package:` field in
// sqlc.yaml is optional — when omitted Go derives it from the import
// path's last segment (here "pgtype"). Setting `package: "pgtype"`
// explicitly triggers a sqlc v1.30 quirk that emits a duplicate alias
// import (`pgtype "github.com/.../pgtype"` alongside the unaliased
// form) and breaks `go build`, so omission is the recommended form.
// This check accepts both.
func isPgtypeOverrideFor(ov sqlcOverrideEntry, dbType, pgPackage, pgType string) bool {
	if !strings.EqualFold(ov.DBType, dbType) {
		return false
	}
	if ov.GoType.Type != pgType {
		return false
	}
	if ov.GoType.Package != "" && ov.GoType.Package != pgPackage {
		return false
	}
	return ov.GoType.Import == "github.com/jackc/pgx/v5/pgtype"
}
