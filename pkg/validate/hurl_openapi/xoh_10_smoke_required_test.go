//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh10SmokeRequired — smoke.hurl 존재 필수 검증

package hurl_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh10SmokeRequired(t *testing.T) {
	cases := []struct {
		name      string
		fs        *yongol.Fullstack
		wantCount int
	}{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{
			name:      "no_hurl_files_produces_diag",
			fs:        &yongol.Fullstack{},
			wantCount: 1,
		},
		{
			name:      "smoke_present_no_diag",
			fs:        &yongol.Fullstack{HurlFiles: []string{"specs/tests/smoke.hurl"}},
			wantCount: 0,
		},
		{
			name:      "smoke_in_subdir",
			fs:        &yongol.Fullstack{HurlFiles: []string{"specs/tests/sub/smoke.hurl"}},
			wantCount: 0,
		},
		{
			name:      "no_smoke_among_files",
			fs:        &yongol.Fullstack{HurlFiles: []string{"specs/tests/login.hurl", "specs/tests/orders.hurl"}},
			wantCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runDiagCodeCase(t, xoh10SmokeRequired(c.fs), c.wantCount, "[XOH-10]")
		})
	}
}
