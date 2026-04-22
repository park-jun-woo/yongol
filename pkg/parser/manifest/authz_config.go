//ff:type feature=projectconfig type=model
//ff:what 인가 패키지 설정 구조체
package manifest

// AuthzConfig configures the authorization package.
type AuthzConfig struct {
	Package string `yaml:"package"` // custom authz package path, default: github.com/park-jun-woo/ssac/pkg/authz
}
