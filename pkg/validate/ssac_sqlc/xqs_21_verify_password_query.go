//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-sqlc
//ff:what xqs21VerifyPasswordQuery — @verify-password 가 참조하는 sqlc 쿼리 존재 강제

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func xqs21VerifyPasswordQuery(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	have := collectHaveQueries(fs)
	var diags []diagnostic.Diagnostic
	for _, f := range fs.ServiceFuncs {
		diags = append(diags, checkVerifyPasswordQueries(f, have)...)
	}
	return diags
}
