//ff:type feature=projectconfig type=model
//ff:what 백엔드 설정 구조체
package manifest

type Backend struct {
	Lang            string                 `yaml:"lang"`
	Framework       string                 `yaml:"framework"`
	Module          string                 `yaml:"module"`
	Middleware      []string               `yaml:"middleware"`
	Auth            *Auth                  `yaml:"auth"`
	CORS            *CORSConfig            `yaml:"cors,omitempty"`
	HTTP            *HTTPConfig            `yaml:"http,omitempty"`
	Observability   *Observability         `yaml:"observability,omitempty"`
	SecurityHeaders *SecurityHeadersConfig `yaml:"security_headers,omitempty"`
	Error           *ErrorConfig           `yaml:"error,omitempty"`
}
