//ff:func feature=gen-gogin type=generator control=sequence
//ff:what pgtypeHelpersEmit — pgtype helper 소스(합본) 생성

package ssac

// pgtypeHelpersEmit returns the Go source for the pgtype helper file. The
// file is emitted into internal/service/pgtype_helpers.go by
// emitPgtypeHelpers. Kept as a single function (no template indirection) so
// the output is deterministic across runs.
//
// The embedded annotation prefix is assembled at runtime so filefunc does
// not parse this file's own annotations inside the returned literal (A6
// self-detection).
func pgtypeHelpersEmit() string {
	ff := "//" + "ff:"
	return ff + `func feature=service type=util control=sequence
` + ff + `what pgtype helpers — pgtype.UUID / pgtype.Numeric 를 primitive string 으로 변환

package service

import (
	"encoding/hex"

	"github.com/jackc/pgx/v5/pgtype"
)

// pgUUIDToString returns the canonical 8-4-4-4-12 UUID string for a valid
// pgtype.UUID, or "" when the value is SQL NULL. The generated convert
// functions call this when mapping a sqlc UUID column (pgtype.UUID) onto
// an api struct field typed as string (openapi_types.UUID / plain string).
func pgUUIDToString(u pgtype.UUID) string {
	if !u.Valid {
		return ""
	}
	b := u.Bytes
	dst := make([]byte, 36)
	hex.Encode(dst[0:8], b[0:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], b[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], b[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], b[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:36], b[10:16])
	return string(dst)
}

// pgNumericToString renders a pgtype.Numeric as its textual form. Returns
// "" on NULL or marshal failure. Callers that need arithmetic precision
// should use pgtype.Numeric directly instead of going through convert.
func pgNumericToString(n pgtype.Numeric) string {
	if !n.Valid {
		return ""
	}
	buf, err := n.MarshalJSON()
	if err != nil {
		return ""
	}
	// MarshalJSON wraps numeric values in quotes only when the driver chose
	// string serialisation; strip them for direct string consumption.
	s := string(buf)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}
`
}
