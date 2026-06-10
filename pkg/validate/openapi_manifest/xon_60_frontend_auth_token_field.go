//ff:func feature=validate type=rule control=sequence topic=config-check
//ff:what XON-60 — frontend.auth 의 token_field/refresh_field/refresh_op 가 OpenAPI 2xx 응답에 실재

package openapi_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xon60FrontendAuthTokenField validates XON-60: when manifest.yaml
// declares frontend.auth, its token_field (and refresh_field when
// declared) must exist as a top-level property of at least one
// operation's 2xx JSON response schema, and refresh_op (when declared)
// must name an existing operationId whose 2xx response carries
// token_field. Without this rule a typo like `token_field: acces_token`
// would generate a capture that silently never matches at runtime —
// the same "manifest claim ↔ source-of-truth" failure mode XDN-* guards
// against (plans/stml/auth-flow Phase001).
func xon60FrontendAuthTokenField(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Frontend.Auth == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	auth := fs.Manifest.Frontend.Auth
	allProps, refreshOpProps, refreshOpFound := collectFrontendAuthProps(fs.OpenAPIDoc, auth.RefreshOp)

	var diags []diagnostic.Diagnostic
	if !allProps[auth.TokenField] {
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XON-60] frontend.auth.token_field \"" + auth.TokenField + "\" not found in any OpenAPI 2xx response schema",
			Advice:  "Pick a field present in a 2xx response schema. " + propCandidates(allProps),
		})
	}
	if auth.RefreshField != "" && !allProps[auth.RefreshField] {
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XON-60] frontend.auth.refresh_field \"" + auth.RefreshField + "\" not found in any OpenAPI 2xx response schema",
			Advice:  "Pick a field present in a 2xx response schema, or remove refresh_field to skip the refresh flow. " + propCandidates(allProps),
		})
	}
	if auth.RefreshOp != "" {
		if !refreshOpFound {
			diags = append(diags, diagnostic.Diagnostic{
				File:    "manifest.yaml",
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[XON-60] frontend.auth.refresh_op \"" + auth.RefreshOp + "\" has no matching OpenAPI operationId",
				Advice:  "Declare the refresh operation in OpenAPI or pick an existing one. " + opCandidates(operationIDSet(fs)),
			})
		} else if auth.TokenField != "" && !refreshOpProps[auth.TokenField] {
			diags = append(diags, diagnostic.Diagnostic{
				File:    "manifest.yaml",
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[XON-60] frontend.auth.refresh_op \"" + auth.RefreshOp + "\" 2xx response has no \"" + auth.TokenField + "\" property — refresh would never yield a new token",
				Advice:  "Add \"" + auth.TokenField + "\" to the refresh operation's 2xx response schema. " + propCandidates(refreshOpProps),
			})
		}
	}
	return diags
}
