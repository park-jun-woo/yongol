//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-cors
//ff:what cors01WildcardCredentials — allow_origins=* + credentials=true 금지 검증
package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCors01WildcardCredentials(t *testing.T) {
	cases := []TestCors01WildcardCredentialsCase{
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "nil_cors", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 0},
		{
			name: "disabled_cors",
			fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
				Backend: pm.Backend{CORS: &pm.CORSConfig{Enabled: false}},
			}},
			wantCount: 0,
		},
		{
			name: "credentials_false",
			fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
				Backend: pm.Backend{CORS: &pm.CORSConfig{Enabled: true, AllowOrigins: []string{"*"}, AllowCredentials: false}},
			}},
			wantCount: 0,
		},
		{
			name: "no_wildcard_with_credentials",
			fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
				Backend: pm.Backend{CORS: &pm.CORSConfig{Enabled: true, AllowOrigins: []string{"https://example.com"}, AllowCredentials: true}},
			}},
			wantCount: 0,
		},
		{
			name: "wildcard_with_credentials_error",
			fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
				Backend: pm.Backend{CORS: &pm.CORSConfig{Enabled: true, AllowOrigins: []string{"*"}, AllowCredentials: true}},
			}},
			wantCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runCors01WildcardCredentials(t, c)
		})
	}
}
