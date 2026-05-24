//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what isPgtypeOverrideFor — sqlc override entry 매칭 (dbType/type/package/import) 전분기 검증

package query

import "testing"

func TestIsPgtypeOverrideFor(t *testing.T) {
	makeEntry := func(dbType, imp, pkg, typ string) sqlcOverrideEntry {
		var e sqlcOverrideEntry
		e.DBType = dbType
		e.GoType.Import = imp
		e.GoType.Package = pkg
		e.GoType.Type = typ
		return e
	}

	const pgxImport = "github.com/jackc/pgx/v5/pgtype"

	t.Run("full match with explicit package", func(t *testing.T) {
		ov := makeEntry("uuid", pgxImport, "pgtype", "UUID")
		if !isPgtypeOverrideFor(ov, "uuid", "pgtype", "UUID") {
			t.Error("expected true")
		}
	})

	t.Run("match with omitted package", func(t *testing.T) {
		ov := makeEntry("uuid", pgxImport, "", "UUID")
		if !isPgtypeOverrideFor(ov, "uuid", "pgtype", "UUID") {
			t.Error("expected true with empty package")
		}
	})

	t.Run("dbType case insensitive", func(t *testing.T) {
		ov := makeEntry("UUID", pgxImport, "", "UUID")
		if !isPgtypeOverrideFor(ov, "uuid", "pgtype", "UUID") {
			t.Error("expected true with uppercase dbType")
		}
	})

	t.Run("wrong dbType", func(t *testing.T) {
		ov := makeEntry("numeric", pgxImport, "", "UUID")
		if isPgtypeOverrideFor(ov, "uuid", "pgtype", "UUID") {
			t.Error("expected false for wrong dbType")
		}
	})

	t.Run("wrong GoType.Type", func(t *testing.T) {
		ov := makeEntry("uuid", pgxImport, "", "Numeric")
		if isPgtypeOverrideFor(ov, "uuid", "pgtype", "UUID") {
			t.Error("expected false for wrong Go type")
		}
	})

	t.Run("wrong package", func(t *testing.T) {
		ov := makeEntry("uuid", pgxImport, "otherpkg", "UUID")
		if isPgtypeOverrideFor(ov, "uuid", "pgtype", "UUID") {
			t.Error("expected false for wrong package")
		}
	})

	t.Run("wrong import", func(t *testing.T) {
		ov := makeEntry("uuid", "github.com/lib/pq", "pgtype", "UUID")
		if isPgtypeOverrideFor(ov, "uuid", "pgtype", "UUID") {
			t.Error("expected false for wrong import")
		}
	})

	t.Run("empty import", func(t *testing.T) {
		ov := makeEntry("uuid", "", "", "UUID")
		if isPgtypeOverrideFor(ov, "uuid", "pgtype", "UUID") {
			t.Error("expected false for empty import")
		}
	})
}
