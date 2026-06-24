//ff:func feature=projectconfig type=util control=sequence
//ff:what DomainConfig.ResolvedAllowOrigins — 도메인 CORS allow_origins 해석 (미지정 시 backend.cors 상속)

package manifest

// ResolvedAllowOrigins returns the origins allowed for this domain. A domain
// that declares its own cors block overrides the backend list wholesale; a
// domain without a cors block inherits backend.cors.allow_origins. Returns nil
// when neither is configured (no cross-origin requests permitted).
func (d DomainConfig) ResolvedAllowOrigins(backend *CORSConfig) []string {
	if d.CORS != nil {
		return d.CORS.AllowOrigins
	}
	if backend != nil {
		return backend.AllowOrigins
	}
	return nil
}
