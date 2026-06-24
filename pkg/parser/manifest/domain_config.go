//ff:type feature=projectconfig type=model
//ff:what DomainConfig — multi-domain 아키텍처 도메인별 설정 (openapi·frontend·route_prefix·auth_mode·cors)
package manifest

// DomainConfig represents one domain entry under the top-level `domains` key.
// A domain is an independent app slice: it owns its OpenAPI contract, its STML
// source directory, the backend route group prefix it mounts under, and may
// override the auth mode and CORS policy it inherits from backend.*.
type DomainConfig struct {
	OpenAPI     string      `yaml:"openapi"`             // domain-specific OpenAPI spec path
	Frontend    string      `yaml:"frontend"`            // domain-specific STML directory
	RoutePrefix string      `yaml:"route_prefix"`        // backend route group prefix (e.g. "/api/admin")
	AuthMode    string      `yaml:"auth_mode,omitempty"` // cookie / bearer / hybrid (inherits backend.auth.mode when omitted)
	CORS        *CORSConfig `yaml:"cors,omitempty"`      // per-domain CORS override (inherits backend.cors when omitted)
}
