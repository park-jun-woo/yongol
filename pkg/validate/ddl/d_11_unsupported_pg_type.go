//ff:func feature=validate type=rule control=iteration dimension=2 topic=ddl-structural
//ff:what D-11 — types.MapPGType.Supported=false 컬럼은 ERROR (다중 토큰 PG 타입 / CREATE TYPE 사용자 ENUM 거절)

package ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d11UnsupportedPgType validates D-11: every parsed DDL column whose
// MapPGType binding has Supported=false is rejected before generate
// runs. The current set of unsupported types covers:
//
//   - multi-word PG type tokens whose single-token alias has no Go-side
//     binding yet (TIME WITH/WITHOUT TIME ZONE → TIMETZ/TIME, BIT
//     VARYING → VARBIT)
//   - user-defined ENUMs declared via CREATE TYPE (yongol does not
//     parse such definitions today)
//
// DOUBLE PRECISION and TIMESTAMP WITH/WITHOUT TIME ZONE used to be
// rejected here, but parser/ddl now preserves them verbatim and
// ddl.NormalizePGTypeHead folds them into FLOAT8 / TIMESTAMPTZ /
// TIMESTAMP, all of which have Go bindings. Inline VARCHAR(N) +
// CHECK IN (...) remains the workaround for enum-like columns.
func d11UnsupportedPgType(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, tbl := range fs.DDLTables {
		for _, col := range tbl.Columns {
			b := types.MapPGType(col)
			if b.Supported {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    tbl.File,
				Line:    tbl.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[D-11] " + tbl.Name + "." + col.Name + " uses unsupported PG type \"" + col.RawType + "\"",
				Advice:  "TIME WITH/WITHOUT TIME ZONE and BIT VARYING have no Go binding yet — choose TIMESTAMPTZ / TIMESTAMP / VARCHAR instead. For CREATE TYPE user ENUMs, use inline VARCHAR(N) + CHECK IN (...) until CREATE TYPE support lands.",
			})
		}
	}
	return diags
}
