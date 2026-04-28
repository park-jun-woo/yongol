//ff:func feature=validate type=rule control=sequence topic=ssac-sqlc
//ff:what buildXqs19Diag — @call / @publish 누락 쿼리 진단 문구 조립

package ssac_sqlc

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func buildXqs19Diag(f ssacparser.ServiceFunc, seq ssacparser.Sequence, pkg, method, query string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:  f.FileName,
		Line:  seq.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XQS-19] %s: @%s %s.%s requires sqlc query %q",
			f.Name, sequenceTag(seq.Type), pkg, method, query),
		Advice: buildXqs19Advice(pkg, query),
	}
}
