//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-21 — bearer 모드인데 auth.token 캡처 0건, 또는 캡처는 있는데 보호 op 호출 페이지 0건 (WARNING)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm21CaptureSinkUnused is the runtime twin of XOH-09 (a captured variable
// must actually be referenced later). In bearer mode it warns when (a) no
// STML page declares an auth.token capture — the generated client would
// never obtain a token — or (b) captures are declared but no page calls a
// security-protected operation, so the captured token is never consumed.
// Cookie/hybrid projects and projects without backend.auth are skipped.
func tm21CaptureSinkUnused(fs *yongol.Fullstack, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	if backendAuthMode(fs) != "bearer" {
		return nil
	}
	captures := collectPageCaptures(fs.STMLPages)
	tokenCaptures := 0
	for _, c := range captures {
		if c.Bind.Sink == "auth.token" {
			tokenCaptures++
		}
	}
	var diags []diagnostic.Diagnostic
	if tokenCaptures == 0 {
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[TM-21] backend.auth.mode is \"bearer\" but no STML page declares an auth.token capture (data-capture)",
			Advice:  "Add data-capture=\"<token_field> -> auth.token\" to the login action block, or switch backend.auth.mode if bearer is not intended",
		})
	}
	if _, ok := firstProtectedOpPage(fs, opMap); len(captures) > 0 && !ok {
		diags = append(diags, diagnostic.Diagnostic{
			File:    captures[0].File,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[TM-21] data-capture declares auth sinks but no STML page calls a security-protected operation — the captured token is never consumed",
			Advice:  "Add a page that calls a protected operation (OpenAPI security), or remove the unused data-capture declaration",
		})
	}
	return diags
}
