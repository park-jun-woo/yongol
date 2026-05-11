//ff:type feature=projectconfig type=model
//ff:what DomainConfig — multi-domain 아키텍처 도메인별 설정 (openapi·frontend 경로)
package manifest

// DomainConfig represents one domain entry under the top-level `domains` key.
type DomainConfig struct {
	OpenAPI  string `yaml:"openapi"`
	Frontend string `yaml:"frontend"`
}
