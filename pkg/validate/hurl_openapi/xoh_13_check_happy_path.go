//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh13CheckHappyPath — @response 함수의 2xx hurl 커버리지 진단

package hurl_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func xoh13CheckHappyPath(fn ssac.ServiceFunc, coveredSet map[string]bool) []diagnostic.Diagnostic {
	var responseLine int
	for _, seq := range fn.Sequences {
		if seq.Type == "response" {
			responseLine = seq.Line
			break
		}
	}
	if responseLine == 0 {
		return nil
	}

	for code := range coveredSet {
		if strings.HasPrefix(code, "2") {
			return nil
		}
	}

	return []diagnostic.Diagnostic{{
		File:    fn.FileName,
		Line:    responseLine,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: fmt.Sprintf("[XOH-13] %s — @response has no 2xx hurl test", fn.Name),
		Advice:  "Add a hurl scenario that gets a 2xx response for this endpoint",
	}}
}
