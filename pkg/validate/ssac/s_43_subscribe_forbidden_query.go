//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-43 — @subscribe에서 query 금지

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s43SubscribeForbiddenQuery validates S-43: @subscribe Args may not use query.
func s43SubscribeForbiddenQuery(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe == nil {
			continue
		}
		for _, seq := range fn.Sequences {
			for _, arg := range seq.Args {
				if arg.Source == "query" {
					diags = append(diags, diagnostic.Diagnostic{
						File:    fn.FileName,
						Line:    seq.Line,
						Phase:   diagnostic.PhaseValidate,
						Level:   diagnostic.LevelError,
						Message: "[S-43] @subscribe cannot use query",
						Advice:  "@subscribe 함수는 HTTP 입력(request/query/currentUser)을 사용할 수 없습니다",
					})
				}
			}
		}
	}
	return diags
}
