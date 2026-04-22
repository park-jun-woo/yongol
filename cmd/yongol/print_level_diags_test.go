//ff:func feature=cli type=test control=iteration dimension=1
//ff:what printLevelDiags 출력 형식 테스트 — file/line prefix 4 케이스
package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestPrintLevelDiagsFileLine(t *testing.T) {
	cases := []struct {
		name string
		diag diagnostic.Diagnostic
		want string
	}{
		{
			name: "file and line",
			diag: diagnostic.Diagnostic{
				File: "service/gig/create.ssac", Line: 42,
				Level: diagnostic.LevelError, Message: "X-99: bad",
			},
			want: "  - service/gig/create.ssac:42: X-99: bad\n",
		},
		{
			name: "file only",
			diag: diagnostic.Diagnostic{
				File: "api/openapi.yaml", Line: 0,
				Level: diagnostic.LevelError, Message: "X-99: bad",
			},
			want: "  - api/openapi.yaml: X-99: bad\n",
		},
		{
			name: "no location",
			diag: diagnostic.Diagnostic{
				File: "", Line: 0,
				Level: diagnostic.LevelError, Message: "X-99: bad",
			},
			want: "  - X-99: bad\n",
		},
		{
			name: "with advice",
			diag: diagnostic.Diagnostic{
				File: "db/users.sql", Line: 5,
				Level: diagnostic.LevelError, Message: "D-2: missing NOT NULL → 권고: 컬럼에 NOT NULL 추가",
			},
			want: "  - db/users.sql:5: D-2: missing NOT NULL\n      ↳ 권고: 컬럼에 NOT NULL 추가\n",
		},
	}
	for _, tc := range cases {
		runPrintLevelDiagsCase(t, tc.name, tc.diag, tc.want)
	}
}
