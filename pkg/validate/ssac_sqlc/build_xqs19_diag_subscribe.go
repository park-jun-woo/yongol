//ff:func feature=validate type=rule control=sequence topic=ssac-sqlc
//ff:what buildXqs19DiagSubscribe — @subscribe 누락 쿼리 진단 문구 조립

package ssac_sqlc

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func buildXqs19DiagSubscribe(f ssacparser.ServiceFunc, query string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:  f.FileName,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XQS-19] %s: @subscribe requires sqlc query %q",
			f.Name, query),
		Advice: buildXqs19Advice("queue", query),
	}
}
