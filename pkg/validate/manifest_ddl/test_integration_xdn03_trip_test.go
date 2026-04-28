//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestIntegration_XDN03_Trip — claims.OrgID: org_id 지만 users 에 org_id 컬럼 없음 → XDN-03

package manifest_ddl

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestIntegration_XDN03_Trip(t *testing.T) {
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
      OrgID: org_id:int64
frontend:
  lang: typescript
  framework: react
  bundler: vite
  name: demo-web
`
	ddlBody := `CREATE TABLE users (
  user_id BIGINT PRIMARY KEY,
  email   VARCHAR(255) NOT NULL
);
`
	root := writeXDNFixture(t, manifestBody, ddlBody)
	cfg, _ := pmanifest.Load(root)
	tables, _ := ddl.ParseTables(filepath.Join(root, "db"))
	fs := &yongol.Fullstack{Manifest: cfg, DDLTables: tables, SpecsDir: root}
	d := Run(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 diagnostic (XDN-03), got %d (%+v)", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDN-03]") || !strings.Contains(d[0].Message, "org_id") {
		t.Fatalf("expected XDN-03 mentioning org_id, got %s", d[0].Message)
	}
}
