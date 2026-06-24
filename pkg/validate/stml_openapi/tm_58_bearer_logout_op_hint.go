//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-58 — bearer 모드 + 값 없는 data-logout + OpenAPI에 logout-like op 존재 시 operationId 명시 유도 (WARNING)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm58BearerLogoutOpHint fires a WARNING when a bearer-mode project has a
// valueless data-logout (OperationID == "") in a layout AND the OpenAPI spec
// contains a logout-like operation (operationId containing "logout",
// case-insensitive, that requires auth). In this configuration the generated
// logout handler only clears the client token and navigates to /login — the
// server-side session (refresh token) is never revoked.
//
// This is the bearer-mode counterpart of TM-38's cookie-mode valueless
// data-logout warning. TM-38 covers "cookie + valueless" and "no auth at
// all"; TM-58 covers "bearer + valueless + a logout op exists in OpenAPI".
func tm58BearerLogoutOpHint(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	auth := prepared.AuthFor(fs)
	if !auth.Present || auth.Mode != "bearer" {
		return nil
	}
	if fs.OpenAPIDoc == nil {
		return nil
	}

	// Find the first logout-like operation that requires auth.
	logoutOp := FindLogoutOp(fs.OpenAPIDoc)
	if logoutOp == "" {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, l := range fs.Layouts {
		if l.Logout == nil {
			continue
		}
		if l.Logout.OperationID != "" {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    l.File,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: fmt.Sprintf("[TM-58] bearer 모드에서 data-logout에 operationId가 미지정되어 서버 logout이 호출되지 않습니다. refresh token이 revoke되지 않을 수 있습니다 (layout %q, candidate op %q)", l.Name, logoutOp),
			Advice:  fmt.Sprintf("data-logout=\"%s\" 로 operationId를 명시하세요", logoutOp),
		})
	}
	return diags
}
