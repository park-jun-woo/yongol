//ff:func feature=projectconfig type=util control=sequence
//ff:what DomainConfig.ResolvedAuthMode — 도메인 auth_mode 해석 (미지정 시 backend 기본값 상속)

package manifest

// ResolvedAuthMode returns the domain's effective auth transport mode. When
// the domain does not declare auth_mode it inherits backendDefault — the
// backend.auth.mode resolved by the caller. The closed set (cookie / bearer /
// hybrid) is validated upstream; this helper only applies the inheritance.
func (d DomainConfig) ResolvedAuthMode(backendDefault string) string {
	if d.AuthMode != "" {
		return d.AuthMode
	}
	return backendDefault
}
