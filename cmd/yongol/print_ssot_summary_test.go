//ff:func feature=cli type=test control=sequence
//ff:what TestPrintSSOTSummary — printSSOTSummary 출력 형식 검증

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestPrintSSOTSummary(t *testing.T) {
	t.Run("EmptyFullstack", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		var buf bytes.Buffer
		printSSOTSummary(&buf, fs)
		out := buf.String()
		if !strings.Contains(out, "SSOT Summary") {
			t.Error("expected 'SSOT Summary' header")
		}
		if !strings.Contains(out, "0 endpoints") {
			t.Error("expected '0 endpoints'")
		}
		if !strings.Contains(out, "0 tables") {
			t.Error("expected '0 tables'")
		}
	})

	t.Run("WithTables", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{{Name: "users"}, {Name: "posts"}},
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "listUsers"},
			},
		}
		var buf bytes.Buffer
		printSSOTSummary(&buf, fs)
		out := buf.String()
		if !strings.Contains(out, "2 tables") {
			t.Errorf("expected '2 tables', got: %s", out)
		}
		if !strings.Contains(out, "1 service functions") {
			t.Errorf("expected '1 service functions', got: %s", out)
		}
	})

	t.Run("WithOpenAPIAndPolicies", func(t *testing.T) {
		paths := openapi3.NewPaths()
		paths.Set("/users", &openapi3.PathItem{
			Get: &openapi3.Operation{OperationID: "listUsers"},
		})
		doc := &openapi3.T{Paths: paths}
		fs := &yongol.Fullstack{
			OpenAPIDoc: doc,
			ParsedPolicies: []rego.Policy{
				{Rules: []rego.AllowRule{{Resource: "user"}, {Resource: "post"}}},
			},
		}
		var buf bytes.Buffer
		printSSOTSummary(&buf, fs)
		out := buf.String()
		if !strings.Contains(out, "1 endpoints") {
			t.Errorf("expected '1 endpoints', got: %s", out)
		}
		if !strings.Contains(out, "2 rules") {
			t.Errorf("expected '2 rules', got: %s", out)
		}
	})
}
