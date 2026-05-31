//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-ddl
//ff:what collectFuncResultDiags — ServiceFunc 의 각 시퀀스에 XDS-12 결과 체크를 적용

package ssac_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func collectFuncResultDiags(fs *yongol.Fullstack, tables map[string]bool, fn ssac.ServiceFunc) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, seq := range fn.Sequences {
		if seq.Type == "call" || seq.Package != "" {
			continue
		}
		if seq.Result == nil || seq.Result.Type == "" {
			continue
		}
		diags = append(diags, checkSeqResultType(fs, tables, fn, seq)...)
	}
	return diags
}
