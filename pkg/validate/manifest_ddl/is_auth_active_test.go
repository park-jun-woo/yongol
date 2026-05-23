//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what isAuthActive — nil/absent/none 비활성 + jwt 등 활성 검증

package manifest_ddl

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestIsAuthActive(t *testing.T) {
	tests := []struct {
		name string
		fs   *yongol.Fullstack
		want bool
	}{
		{
			name: "nil fullstack",
			fs:   nil,
			want: false,
		},
		{
			name: "nil manifest",
			fs:   &yongol.Fullstack{},
			want: false,
		},
		{
			name: "nil auth",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{},
			},
			want: false,
		},
		{
			name: "auth type none",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: "none"},
					},
				},
			},
			want: false,
		},
		{
			name: "auth type jwt is active",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: "jwt"},
					},
				},
			},
			want: true,
		},
		{
			name: "auth type empty string is active (default)",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: ""},
					},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAuthActive(tt.fs)
			if got != tt.want {
				t.Errorf("isAuthActive() = %v, want %v", got, tt.want)
			}
		})
	}
}
