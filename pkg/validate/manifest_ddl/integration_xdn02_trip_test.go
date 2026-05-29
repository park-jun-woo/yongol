//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestIntegration_XDN02_Trip — user_table=users 인데 db/users.sql 부재 → XDN-02 ERROR

package manifest_ddl

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestIntegration_XDN02_Trip(t *testing.T) {
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
      Email: email:string
frontend:
  lang: typescript
  framework: react
  bundler: vite
  name: demo-web
`
	root := writeXDNFixture(t, manifestBody, "")
	cfg, parseDiags := pmanifest.Load(root)
	if len(parseDiags) != 0 {
		t.Fatalf("manifest parse diags: %+v", parseDiags)
	}
	fs := &yongol.Fullstack{Manifest: cfg, SpecsDir: root}
	d := Run(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 diagnostic (XDN-02), got %d (%+v)", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDN-02]") {
		t.Fatalf("expected XDN-02, got %s", d[0].Message)
	}
}
