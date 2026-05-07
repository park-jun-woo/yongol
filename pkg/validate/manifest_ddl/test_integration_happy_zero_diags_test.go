//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestIntegration_Happy_ZeroDiags — manifest+users.sql 정합 시 모든 규칙 진단 0

package manifest_ddl

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestIntegration_Happy_ZeroDiags(t *testing.T) {
	manifestBody := `apiVersion: yongol/v1
kind: Project
metadata:
  name: demo
backend:
  lang: go
  framework: gin
  module: github.com/example/demo
  auth:
    type: jwt
    secret_env: JWT_SECRET
    user_table: users
    claims:
      ID: user_id:int64
      Email: email:string
      Role: role:string
frontend:
  lang: typescript
  framework: react
  bundler: vite
  name: demo-web
`
	ddlBody := `CREATE TABLE users (
  user_id BIGINT PRIMARY KEY,
  email   VARCHAR(255) NOT NULL,
  role    VARCHAR(32)  NOT NULL
);
`
	root := writeXDNFixture(t, manifestBody, ddlBody)

	cfg, parseDiags := pmanifest.Load(root)
	if len(parseDiags) != 0 {
		t.Fatalf("manifest parse diags: %+v", parseDiags)
	}
	tables, ddlDiags := ddl.ParseTables(filepath.Join(root, "db"))
	if len(ddlDiags) != 0 {
		t.Fatalf("ddl parse diags: %+v", ddlDiags)
	}
	fs := &yongol.Fullstack{Manifest: cfg, DDLTables: tables, SpecsDir: root}
	if d := Run(fs); len(d) != 0 {
		t.Fatalf("expected 0 diagnostics, got %+v", d)
	}
}
