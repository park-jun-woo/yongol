//ff:func feature=gen-gogin type=util control=sequence
//ff:what wrapJSONBLiteral — sqlc INSERT 의 JSONB 컬럼 자리에 string 리터럴이 오면 []byte(...) 로 래핑 (BUG-037 #1)

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
)

// wrapJSONBLiteral applies the JSONB literal → []byte wrap on
// expressions that the SSaC author wrote as a Go string literal but
// land in a sqlc params field whose underlying column is JSONB / JSON.
//
// Example:
//
//	DDL                       :  claims JSONB NOT NULL DEFAULT '{}'
//	SSaC                      :  Claims: "{}"
//	without wrap              :  db.UserCreateParams{Claims: "{}"}   // type error: []byte expects []byte not string
//	with wrap (this helper)   :  db.UserCreateParams{Claims: []byte("{}")}
//
// The wrap fires only when:
//   - the method name resolves to a DDL table via the sqlc query catalogue
//   - that table has a column matching the input key (PascalCase ↔ lower)
//   - the column's MapPGType binding is KindJSONB
//   - the rendered expression is a Go string literal ("..." form)
//
// Any other case returns rendered unchanged so existing flows keep
// working. The function is intentionally narrow because only string
// literals miscompile — variables / expressions of type string would
// already be flagged by validate or would produce a different error.
func (g *methodGen) wrapJSONBLiteral(inputKey, rendered string) string {
	if rendered == "" || !looksLikeStringLiteral(rendered) {
		return rendered
	}
	col := g.lookupSQLCMethodColumn(inputKey)
	if col == nil {
		return rendered
	}
	binding := types.MapPGType(*col)
	if binding.Kind != types.KindJSONB {
		return rendered
	}
	return "[]byte(" + rendered + ")"
}

// looksLikeStringLiteral checks whether s is a Go double-quoted string
// literal (no escaping inspection — the wrap is byte-level).
func looksLikeStringLiteral(s string) bool {
	return len(s) >= 2 && strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`)
}
