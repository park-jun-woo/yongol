//ff:type feature=projectconfig type=model
//ff:what HSTSConfig — backend.security_headers.hsts 모델 (Strict-Transport-Security 설정)

package manifest

// HSTSConfig controls the Strict-Transport-Security header. MaxAge is seconds.
// Preload requires MaxAge >= 31536000 and IncludeSubDomains=true to be
// accepted by browser preload lists; yongol does not enforce this — the
// SEC-302 validate rule warns when MaxAge is too short.
type HSTSConfig struct {
	MaxAge            int  `yaml:"max_age,omitempty"`
	IncludeSubDomains bool `yaml:"include_subdomains,omitempty"`
	Preload           bool `yaml:"preload,omitempty"`
}
