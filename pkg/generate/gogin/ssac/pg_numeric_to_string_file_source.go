//ff:func feature=gen-gogin type=generator control=sequence
//ff:what pgNumericToStringFileSource — pgNumericToString 헬퍼의 단일 파일 Go 소스

package ssac

// pgNumericToStringFileSource returns the full Go source of the
// pgNumericToString helper in its own file (F1 compliant). The ff prefix
// is assembled at runtime so filefunc does not parse the embedded
// annotations as this file's own (A6 self-detection).
func pgNumericToStringFileSource() string {
	ff := "//" + "ff:"
	return ff + `func feature=service type=util control=sequence
` + ff + `what pgNumericToString — pgtype.Numeric → textual form (null / 실패 는 "")

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
`
}
