//ff:type feature=gen-gogin type=model
//ff:what resolvedSecurityHeaders — manifest + default 를 병합한 보안 헤더 generator-time 뷰

package boot

// resolvedSecurityHeaders captures the generator-time view of
// manifest.backend.security_headers after defaults are applied. Runtime env
// overrides are layered on top in the generated main.go.
type resolvedSecurityHeaders struct {
	Profile           string
	HSTSMaxAge        int
	HSTSIncludeSubs   bool
	HSTSPreload       bool
	CSPEnabled        bool
	CSPReportOnly     bool
	CSPDirectives     map[string][]string
	XFrameOptions     string
	ReferrerPolicy    string
	PermissionsPolicy map[string][]string
}
