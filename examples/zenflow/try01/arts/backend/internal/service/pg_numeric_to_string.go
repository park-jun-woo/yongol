//ff:func feature=service type=util control=sequence
//ff:what pgNumericToString — pgtype.Numeric → textual form (null / 실패 는 "")
//ff:checked llm=yongol-gen hash=5da1857f

package service

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// pgNumericToString renders a pgtype.Numeric as its textual form. Returns
// "" on NULL or marshal failure. Callers that need arithmetic precision
// should consume pgtype.Numeric directly instead of going through convert.
func pgNumericToString(n pgtype.Numeric) string {
	if !n.Valid {
		return ""
	}
	buf, err := n.MarshalJSON()
	if err != nil {
		return ""
	}
	s := string(buf)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}
