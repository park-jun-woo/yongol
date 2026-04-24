//ff:func feature=gen-gogin type=generator control=sequence topic=pgtype-unwrap
//ff:what emitPgtypeHelpers — internal/service/pgtype_helpers.go 작성 (pgUUIDToString 등)

package ssac

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/fffile"
)

// emitPgtypeHelpers writes internal/service/pgtype_helpers.go holding the
// pgtype → primitive conversion helpers referenced by generated convert
// functions (pgUUIDToString, pgNumericToString). Returning two funcs in a
// single file violates filefunc F1, so the helpers live in separate files
// split out of a shared source; this emitter keeps them co-located for
// now and leaves the F1 split to a follow-up (same pattern as
// emit_convert_func_file.go).
func emitPgtypeHelpers(serviceDir string) error {
	// Emit pgUUIDToString and pgNumericToString as separate files so F1
	// (1-file-1-func) is honoured out of the gate.
	uuid := pgUUIDToStringFileSource()
	numeric := pgNumericToStringFileSource()
	if err := fffile.WriteIfNotPreserved(filepath.Join(serviceDir, "pg_uuid_to_string.go"), []byte(uuid)); err != nil {
		return err
	}
	if err := fffile.WriteIfNotPreserved(filepath.Join(serviceDir, "pg_numeric_to_string.go"), []byte(numeric)); err != nil {
		return err
	}
	return nil
}

// pgUUIDToStringFileSource returns the full Go source of the
// pgUUIDToString helper in its own file (F1 compliant).
func pgUUIDToStringFileSource() string {
	return `//ff:func feature=service type=util control=sequence topic=pgtype-unwrap
//ff:what pgUUIDToString — pgtype.UUID → canonical string (null 은 "")

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

// pgNumericToStringFileSource returns the full Go source of the
// pgNumericToString helper in its own file (F1 compliant).
func pgNumericToStringFileSource() string {
	return `//ff:func feature=service type=util control=sequence topic=pgtype-unwrap
//ff:what pgNumericToString — pgtype.Numeric → textual form (null / 실패 는 "")

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
