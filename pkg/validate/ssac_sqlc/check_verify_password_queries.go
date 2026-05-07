//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what checkVerifyPasswordQueries — 단일 ServiceFunc 의 @verify-password 쿼리 존재 검사

package ssac_sqlc

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func checkVerifyPasswordQueries(f ssacparser.ServiceFunc, have map[string]bool) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, seq := range f.Sequences {
		if seq.Type != "verify-password" {
			continue
		}
		queryName := seq.Model + "FindBy" + pascalCase(seq.EmailCol)
		if have[queryName] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[XQS-21] @verify-password in %s requires sqlc query %q, but it does not exist", f.Name, queryName),
			Advice:  fmt.Sprintf("Add a :one query named %q to db/%s.sql", queryName, modelToTableName(seq.Model)),
		})
	}
	return diags
}
