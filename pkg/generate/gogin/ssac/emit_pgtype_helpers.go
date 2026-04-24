//ff:func feature=gen-gogin type=generator control=sequence
//ff:what emitPgtypeHelpers — internal/service/pgtype_helpers.go 작성 (pgUUIDToString 등)

package ssac

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/fffile"
)

// emitPgtypeHelpers writes internal/service/pgtype_helpers.go holding the
// pgtype → primitive conversion helpers referenced by generated convert
// functions (pgUUIDToString, pgNumericToString). Each helper lives in its
// own file (pgUUIDToStringFileSource / pgNumericToStringFileSource) so the
// emitted sources honour filefunc F1.
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
