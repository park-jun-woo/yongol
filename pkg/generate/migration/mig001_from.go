//ff:func feature=migration type=util control=sequence
//ff:what mig001From — generate 파이프라인에서 MIG-001 rename 불일치 진단 (validate 패키지 import 회피)
package migration

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// mig001From is a generate-side mirror of validate/migration.Mig001RenameMismatch.
// Avoids the validate -> generate import cycle by reimplementing the
// small check here. Intentional duplication: the validate package is
// the normative owner; this version runs inside the generate pipeline.
func mig001From(prev, curr *Schema, hints *Hints) []diagnostic.Diagnostic {
	if hints == nil {
		return nil
	}
	var out []diagnostic.Diagnostic
	out = append(out, mig001CheckRenameTables(prev, curr, hints.RenameTables)...)
	out = append(out, mig001CheckRenameColumns(prev, curr, hints.RenameColumns)...)
	return out
}
