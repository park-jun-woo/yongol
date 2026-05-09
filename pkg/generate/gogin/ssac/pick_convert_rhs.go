//ff:func feature=gen-gogin type=util control=selection
//ff:what pickConvertRHS — convert<X> struct literal 의 오른쪽 표현식 선택 (JSONB/types.Expand/cast/required 분기)

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// pickConvertRHS chooses the right-hand side for one struct-literal line.
// JSONB fields read from the pre-unmarshalled local map variable;
// api-typed fields (enum or format-string wrappers like
// openapi_types.Email) cast through the api-side named type since sqlc
// keeps these columns as plain string; required scalars assign the
// converted row value directly; optional scalars wrap with ptrOf so the
// *T api slot accepts them.
//
// Phase001 (types/ unification) — when col != nil the row → model
// expression is sourced from types.MapPGType(col).ConvertExpr expanded
// with {row}=row, {field}=dbField. Pgtype unwrap (pgtypex.FromPg*)
// is encoded in the ConvertExpr template so callers no longer maintain
// a parallel kind ladder. col == nil falls back to the historic
// `row.<dbField>` direct-assign path used by api wrappers without a
// backing DDL column.
func pickConvertRHS(jsonName, apiField, dbField string, isRequired bool, jsonbs []jsonbFieldAlias, apiCast string, col *ddl.Column) string {
	if rhs, ok := jsonbConvertRHS(jsonName, isRequired, jsonbs); ok {
		return rhs
	}
	base := convertBaseExpr(dbField, col)
	switch {
	case apiCast != "" && isRequired:
		return apiCast + "(" + base + ")"
	case apiCast != "":
		return "ptrOf(" + apiCast + "(" + base + "))"
	case isRequired:
		return base
	default:
		if col != nil {
			binding := types.MapPGType(*col)
			if strings.HasPrefix(binding.ApiField, "*") {
				return base
			}
		}
		return "ptrOf(" + base + ")"
	}
}
