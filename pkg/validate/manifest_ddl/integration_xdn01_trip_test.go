//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestIntegration_XDN01_Trip — auth.type=jwt 인데 user_table 누락 → XDN-01 ERROR

package manifest_ddl

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestIntegration_XDN01_Trip(t *testing.T) {
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
	if len(d) == 0 {
		t.Fatalf("expected XDN-01 trip; got 0 diagnostics")
	}
	if !strings.Contains(d[0].Message, "[XDN-01]") {
		t.Fatalf("first diag should be XDN-01: %s", d[0].Message)
	}
}
