//ff:func feature=gen-gogin type=generator control=sequence
//ff:what pgUUIDToStringFileSource — pgUUIDToString 헬퍼의 단일 파일 Go 소스

package ssac

// pgUUIDToStringFileSource returns the full Go source of the
// pgUUIDToString helper in its own file (F1 compliant).
//
// The returned text starts with //ff annotations that must end up inside
// the generated file but must NOT be parsed as THIS file's own
// annotations. We build the ff prefix at runtime via string concatenation
// so filefunc's lexer sees only broken fragments when scanning this
// source file (filefunc A6 self-detection).
func pgUUIDToStringFileSource() string {
	ff := "//" + "ff:"
	return ff + `func feature=service type=util control=sequence
` + ff + `what pgUUIDToString — pgtype.UUID → canonical string (null 은 "")

package service

import (
	"encoding/hex"

	"github.com/jackc/pgx/v5/pgtype"
)

// pgUUIDToString returns the canonical 8-4-4-4-12 UUID string for a valid
// pgtype.UUID, or "" when the value is SQL NULL. Generated convert
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
`
}
