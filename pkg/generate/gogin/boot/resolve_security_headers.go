//ff:func feature=gen-gogin type=util control=sequence topic=security-headers
//ff:what resolveSecurityHeaders — manifest.security_headers + production 기본값 병합

package boot

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// resolveSecurityHeaders collapses manifest config with production-profile
// defaults. A nil block yields the full production preset. Individual nil /
// zero sub-blocks fall back to their own defaults.
func resolveSecurityHeaders(fs *yongol.Fullstack) resolvedSecurityHeaders {
	out := resolvedSecurityHeaders{
		Profile:         defaultSHProfile,
		HSTSMaxAge:      defaultHSTSMaxAge,
		HSTSIncludeSubs: defaultHSTSIncludeSubs,
		HSTSPreload:     defaultHSTSPreload,
		CSPEnabled:      true,
		CSPReportOnly:   false,
		CSPDirectives: map[string][]string{
			"default-src":     {"'self'"},
			"frame-ancestors": {"'none'"},
			"base-uri":        {"'self'"},
		},
		XFrameOptions:  defaultXFrameOptions,
		ReferrerPolicy: defaultReferrerPolicy,
		PermissionsPolicy: map[string][]string{
			"camera":      {},
			"microphone":  {},
			"geolocation": {},
		},
	}

	if fs == nil || fs.Manifest == nil {
		return out
	}
	sh := fs.Manifest.Backend.SecurityHeaders
	if sh == nil {
		return out
	}

	if p := strings.TrimSpace(sh.Profile); p != "" {
		out.Profile = strings.ToLower(p)
	}
	if sh.HSTS != nil {
		if sh.HSTS.MaxAge > 0 {
			out.HSTSMaxAge = sh.HSTS.MaxAge
		}
		out.HSTSIncludeSubs = sh.HSTS.IncludeSubDomains
		out.HSTSPreload = sh.HSTS.Preload
	}
	if sh.CSP != nil {
		if sh.CSP.Enabled != nil {
			out.CSPEnabled = *sh.CSP.Enabled
		}
		out.CSPReportOnly = sh.CSP.ReportOnly
		if len(sh.CSP.Directives) > 0 {
			out.CSPDirectives = copyStringListMap(sh.CSP.Directives)
		}
	}
	if v := strings.TrimSpace(sh.XFrameOptions); v != "" {
		out.XFrameOptions = v
	}
	if v := strings.TrimSpace(sh.ReferrerPolicy); v != "" {
		out.ReferrerPolicy = v
	}
	if sh.PermissionsPolicy != nil {
		out.PermissionsPolicy = copyStringListMap(sh.PermissionsPolicy)
	}
	return out
}
