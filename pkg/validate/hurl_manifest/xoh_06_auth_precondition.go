//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-manifest
//ff:what XOH-06 — 보호 구간 operation 호출 전에 인증 스텝이 선행되어야 함

package hurl_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh06AuthPrecondition enforces XOH-06 (WARNING): when the manifest
// declares an auth middleware, any hurl call to a protected endpoint
// must be preceded — within the same file — by either an auth-issuing
// step or a request carrying the auth credential.
func xoh06AuthPrecondition(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return nil
	}
	byFile := groupByFile(fs.HurlEntries)
	var diags []diagnostic.Diagnostic
	for _, entries := range byFile {
		diags = append(diags, checkFileAuth(entries)...)
	}
	return diags
}
