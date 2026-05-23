//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh09CheckEntryCaptures — capture 변수가 파일 내에서 재사용되는지 확인 검증

package hurl_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestXoh09CheckEntryCaptures(t *testing.T) {
	cases := []struct {
		name      string
		file      string
		text      string
		entry     hurl.HurlEntry
		wantCount int
	}{
		{
			name:      "no_captures",
			file:      "t.hurl",
			text:      "",
			entry:     hurl.HurlEntry{},
			wantCount: 0,
		},
		{
			name: "used_capture_no_diag",
			file: "t.hurl",
			text: "GET /api/users\nAuthorization: Bearer {{tok}}",
			entry: hurl.HurlEntry{
				Captures: []hurl.HurlCapture{{Var: "tok", Line: 3}},
			},
			wantCount: 0,
		},
		{
			name: "unused_capture_produces_warning",
			file: "t.hurl",
			text: "GET /api/users\n",
			entry: hurl.HurlEntry{
				Captures: []hurl.HurlCapture{{Var: "tok", Line: 3}},
			},
			wantCount: 1,
		},
		{
			name: "empty_var_name_skipped",
			file: "t.hurl",
			text: "GET /api/users\n",
			entry: hurl.HurlEntry{
				Captures: []hurl.HurlCapture{{Var: "", Line: 3}},
			},
			wantCount: 0,
		},
		{
			name: "mixed_used_and_unused",
			file: "t.hurl",
			text: "GET /api/users\nAuthorization: Bearer {{tok}}",
			entry: hurl.HurlEntry{
				Captures: []hurl.HurlCapture{
					{Var: "tok", Line: 3},
					{Var: "unused_var", Line: 4},
				},
			},
			wantCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runDiagCodeCase(t, xoh09CheckEntryCaptures(c.file, c.text, c.entry), c.wantCount, "[XOH-09]")
		})
	}
}
