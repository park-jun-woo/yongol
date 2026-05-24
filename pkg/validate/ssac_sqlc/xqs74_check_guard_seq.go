//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what xqs74CheckGuardSeq — 단일 @empty/@exists 시퀀스의 PK 타입 제약 위반 검사

package ssac_sqlc

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xqs74CheckGuardSeq checks a single @empty/@exists sequence against the PK
// type constraint. Returns (diag, true) on violation.
func xqs74CheckGuardSeq(seq ssacparser.Sequence, varModel map[string]string, tableMap map[string]*ddl.Table, fileName string) (diagnostic.Diagnostic, bool) {
	if seq.Type != "empty" && seq.Type != "exists" {
		return diagnostic.Diagnostic{}, false
	}
	target := seq.Target
	if target == "" || strings.Contains(target, ".") {
		return diagnostic.Diagnostic{}, false
	}
	modelName, ok := varModel[target]
	if !ok {
		return diagnostic.Diagnostic{}, false
	}
	return xqs74CheckModel(seq, modelName, tableMap, fileName)
}
