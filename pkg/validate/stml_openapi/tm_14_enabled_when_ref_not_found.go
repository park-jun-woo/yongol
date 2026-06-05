//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-14 — data-enabled-when 가드 ref의 model 접두어가 페이지 fetch 응답 스키마에 없음 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm14EnabledWhenRefNotFound checks that every model prefix referenced by an
// action's data-enabled-when guard is a top-level property of some page fetch
// response (modelFetchMap, built by buildModelFetchMap). The scope matches
// TM-06: only the top-level model key is checked, not nested fields. Guard
// syntax errors are handled earlier by TM-17, so a parse failure here is
// silently skipped. An empty EnabledWhen yields no diagnostics.
func tm14EnabledWhenRefNotFound(a stml.ActionBlock, file string, modelFetchMap map[string]operationEntry) []diagnostic.Diagnostic {
	if a.EnabledWhen == "" {
		return nil
	}
	expr, err := stml.ParseGuard(a.EnabledWhen)
	if err != nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, ref := range expr.CollectRefs() {
		if _, ok := modelFetchMap[ref.Model]; ok {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    file,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-14] data-enabled-when ref %q on action %q references model %q, which is not in any page fetch response schema", ref.Path(), a.OperationID, ref.Model),
			Advice:  fmt.Sprintf("Add a data-fetch on this page whose response has top-level property %q, or change the data-enabled-when guard to reference a fetched resource", ref.Model),
		})
	}
	return diags
}
