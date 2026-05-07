//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestIntegration_XDN06_Trip — claims.IsAdmin: is_admin:int64 인데 컬럼이 BOOLEAN → XDN-06

package manifest_ddl

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestIntegration_XDN06_Trip(t *testing.T) {
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
      IsAdmin: is_admin:int64
frontend:
  lang: typescript
  framework: react
  bundler: vite
  name: demo-web
`
	ddlBody := `CREATE TABLE users (
  user_id  BIGINT  PRIMARY KEY,
  is_admin BOOLEAN NOT NULL DEFAULT FALSE
);
`
	root := writeXDNFixture(t, manifestBody, ddlBody)
	cfg, _ := pmanifest.Load(root)
	tables, _ := ddl.ParseTables(filepath.Join(root, "db"))
	fs := &yongol.Fullstack{Manifest: cfg, DDLTables: tables, SpecsDir: root}
	d := Run(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 diagnostic (XDN-06), got %d (%+v)", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDN-06]") {
		t.Fatalf("expected XDN-06, got %s", d[0].Message)
	}
}
