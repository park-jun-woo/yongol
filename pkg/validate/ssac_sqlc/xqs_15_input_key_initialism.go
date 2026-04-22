//ff:func feature=validate type=rule control=iteration dimension=3 topic=sqlc
//ff:what XQS-15 — SSaC input key 가 Go initialism 컨벤션을 위반하는지 검사

package ssac_sqlc

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// xqs15InputKeyInitialism flags SSaC input keys that violate Go initialism
// conventions (Id → ID, OrgId → OrgID, Url → URL, ...). Go struct fields use
// the corrected form; generated code with mismatched case breaks compilation
// or silently shadows the correct field.
//
// Complementary to XQS-14 (case mismatch vs sqlc param). XQS-14 accepts
// snake-case bridging (`Email` ↔ `email`); XQS-15 catches the specific
// case where the PascalCase key itself is malformed.
func xqs15InputKeyInitialism(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	seen := map[string]bool{}
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			// CRUD sequences use sqlc params — sqlc has its own casing rules (XQS-14/16).
			// XQS-15 only applies to @call sequences.
			if seq.Type != "call" {
				continue
			}
			keys := collectInputKeys(seq)
			for _, key := range keys {
				correct, violates := rule.ViolatesInitialism(key)
				if !violates {
					continue
				}
				dedupKey := fn.FileName + "|" + fn.Name + "|" + key
				if seen[dedupKey] {
					continue
				}
				seen[dedupKey] = true
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelWarning,
					Message: fmt.Sprintf("[XQS-15] input key %q violates Go initialism (should be %q)", key, correct),
					Advice:  "Go struct field 로 코드젠되므로 약자는 모두 대문자로 표기하세요",
				})
			}
		}
	}
	return diags
}
