//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-22 — bearer 모드에서 보호 op 호출 페이지가 있는데 auth.token 캡처가 전무 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm22ProtectedOpNoTokenSupply is the runtime twin of XOH-06 (a protected
// call must be preceded by an auth step). In bearer mode, when at least one
// STML page calls a security-protected operation but no page anywhere
// captures auth.token, every protected screen is guaranteed a 401 — the
// generated client has no token source. That contradiction is an ERROR.
func tm22ProtectedOpNoTokenSupply(fs *yongol.Fullstack, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	if backendAuthMode(fs) != "bearer" {
		return nil
	}
	file, ok := firstProtectedOpPage(fs, opMap)
	if !ok {
		return nil
	}
	for _, c := range collectPageCaptures(fs.STMLPages) {
		if c.Bind.Sink == "auth.token" {
			return nil
		}
	}
	return []diagnostic.Diagnostic{{
		File:    file,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[TM-22] bearer mode: %q calls a security-protected operation but no STML page captures auth.token — every protected screen is guaranteed a 401", file),
		Advice:  "Add data-capture=\"<token_field> -> auth.token\" to the login action block so the generated client can supply the Bearer token",
	}}
}
