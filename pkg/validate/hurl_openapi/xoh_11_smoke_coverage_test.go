//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh11SmokeCoverage — smoke.hurl이 모든 operationId를 커버하는지 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh11SmokeCoverage(t *testing.T) {
	doc := &openapi3.T{Paths: &openapi3.Paths{}}
	op1 := &openapi3.Operation{OperationID: "getUsers", Responses: &openapi3.Responses{}}
	op2 := &openapi3.Operation{OperationID: "createUser", Responses: &openapi3.Responses{}}
	doc.Paths.Set("/users", &openapi3.PathItem{Get: op1, Post: op2})

	cases := []struct {
		name      string
		fs        *yongol.Fullstack
		wantCount int
	}{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_doc", fs: &yongol.Fullstack{}, wantCount: 0},
		{
			name: "no_smoke_file_no_diag",
			fs: &yongol.Fullstack{
				OpenAPIDoc: doc,
				HurlFiles:  []string{"specs/tests/login.hurl"},
			},
			wantCount: 0,
		},
		{
			name: "full_coverage_no_diag",
			fs: &yongol.Fullstack{
				OpenAPIDoc: doc,
				HurlFiles:  []string{"specs/tests/smoke.hurl"},
				HurlEntries: []hurl.HurlEntry{
					{Method: "GET", Path: "/users", File: "specs/tests/smoke.hurl"},
					{Method: "POST", Path: "/users", File: "specs/tests/smoke.hurl"},
				},
			},
			wantCount: 0,
		},
		{
			name: "no_operation_ids_no_diag",
			fs: func() *yongol.Fullstack {
				d := &openapi3.T{Paths: &openapi3.Paths{}}
				d.Paths.Set("/health", &openapi3.PathItem{
					Get: &openapi3.Operation{Responses: &openapi3.Responses{}},
				})
				return &yongol.Fullstack{
					OpenAPIDoc: d,
					HurlFiles:  []string{"specs/tests/smoke.hurl"},
				}
			}(),
			wantCount: 0,
		},
		{
			name: "partial_coverage_produces_diag",
			fs: &yongol.Fullstack{
				OpenAPIDoc: doc,
				HurlFiles:  []string{"specs/tests/smoke.hurl"},
				HurlEntries: []hurl.HurlEntry{
					{Method: "GET", Path: "/users", File: "specs/tests/smoke.hurl"},
				},
			},
			wantCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runDiagCodeCase(t, xoh11SmokeCoverage(c.fs), c.wantCount, "[XOH-11]")
		})
	}
}
