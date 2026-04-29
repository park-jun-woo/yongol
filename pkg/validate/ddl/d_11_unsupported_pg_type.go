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
// runs. The current set of unsupported types covers multi-word PG type
// tokens that the parser preserves verbatim ("DOUBLE PRECISION",
// "TIMESTAMP WITH TIME ZONE") and user-defined ENUMs declared via
// CREATE TYPE (yongol does not parse such definitions today).
//
// Suggested replacements: use TIMESTAMPTZ instead of "TIMESTAMP WITH
// TIME ZONE", FLOAT8 instead of "DOUBLE PRECISION", and inline
// VARCHAR(N) + CHECK IN (...) for enum-like columns until CREATE TYPE
// support is added.
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
				Advice:  "Use a single-token alias (TIMESTAMPTZ for TIMESTAMP WITH TIME ZONE, FLOAT8 for DOUBLE PRECISION) or wait for CREATE TYPE / multi-token PG type support. Inline VARCHAR(N) + CHECK IN (...) covers enum-like cases today.",
			})
		}
	}
	return diags
}
