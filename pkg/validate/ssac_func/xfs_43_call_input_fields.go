//ff:func feature=validate type=rule control=iteration dimension=3 topic=func-check
//ff:what XFS-43 — @call Input field → FuncRequest

package ssac_func

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xfs43CallInputFields validates XFS-43: @call Input field → FuncRequest
func xfs43CallInputFields(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "call" || seq.Model == "" {
				continue
			}
			funcName := callFuncName(seq.Model)
			camelName := callFuncCamelName(seq.Model)
			if funcName == "" || camelName == "" {
				continue
			}
			reqFields, ok := g.Schemas["Func.request."+camelName]
			if !ok {
				continue
			}
			reqSet := make(map[string]bool, len(reqFields))
			for _, f := range reqFields {
				reqSet[f] = true
			}
			for inputKey := range seq.Inputs {
				if !reqSet[inputKey] {
					diags = append(diags, diagnostic.Diagnostic{
						File:    fn.FileName,
						Line:    seq.Line,
						Phase:   diagnostic.PhaseValidate,
						Level:   diagnostic.LevelError,
						Message: "[XFS-43] @call " + seq.Model + " input field " + inputKey + " not in " + funcName + "Request",
						Advice:  "@call 입력 키 " + inputKey + " 를 func Request struct 에 추가하거나 SSaC 에서 제거하세요",
					})
				}
			}
		}
	}
	return diags
}
