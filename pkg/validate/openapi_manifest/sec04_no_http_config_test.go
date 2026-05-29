//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what SEC-04 — backend.http 가 없으면 규칙 비활성

package openapi_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSEC04_NoHTTPConfig(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: buildDocWithOps([]string{"Any"}),
		Manifest:   &manifest.ProjectConfig{},
	}
	if diags := sec04HTTPOverridesOperationID(fs); diags != nil {
		t.Fatalf("expected nil diags, got %+v", diags)
	}
}
