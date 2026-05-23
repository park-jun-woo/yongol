//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what xnq90QueueBackendRequiresSQLC — nil/absent/memory 조기 반환 + postgres 시 위임 검증

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnq90QueueBackendRequiresSQLC(t *testing.T) {
	tests := []struct {
		name      string
		fs        *yongol.Fullstack
		wantDiags bool
		wantSub   string
	}{
		{
			name:      "nil fullstack returns nil",
			fs:        nil,
			wantDiags: false,
		},
		{
			name:      "nil manifest returns nil",
			fs:        &yongol.Fullstack{},
			wantDiags: false,
		},
		{
			name: "nil queue returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{},
			},
			wantDiags: false,
		},
		{
			name: "memory backend returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Queue: &pmanifest.QueueBackend{Backend: "memory"},
				},
			},
			wantDiags: false,
		},
		{
			name: "postgres backend with no DDL/sqlc raises diagnostic",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Queue: &pmanifest.QueueBackend{Backend: "postgres"},
				},
			},
			wantDiags: true,
			wantSub:   "XNQ-90",
		},
		{
			name: "diagnostic mentions fullend_queue",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Queue: &pmanifest.QueueBackend{Backend: "postgres"},
				},
			},
			wantDiags: true,
			wantSub:   "fullend_queue",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := xnq90QueueBackendRequiresSQLC(tt.fs)
			assertDiags(t, diags, tt.wantDiags, tt.wantSub)
		})
	}
}
