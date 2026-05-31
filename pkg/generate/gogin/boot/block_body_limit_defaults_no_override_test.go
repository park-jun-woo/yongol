//ff:func feature=gen-gogin type=test control=sequence topic=dos-guard
//ff:what blockBodyLimit — middleware.BodyLimit / MultipartLimit / OverrideBodyLimit 등록
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockBodyLimit_DefaultsNoOverride(t *testing.T) {
	block := blockBodyLimit(&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `envInt64("BACKEND_HTTP_BODY_LIMIT", 1048576)`) {
		t.Errorf("default body limit (1MiB) wrong, got:\n%s", body)
	}
	if !strings.Contains(body, "r.Use(middleware.BodyLimit(bodyLimit))") {
		t.Errorf("must register BodyLimit, got:\n%s", body)
	}
	if strings.Contains(body, "OverrideBodyLimit") {
		t.Errorf("no override should be emitted without manifest overrides, got:\n%s", body)
	}
}
