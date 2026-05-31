//ff:func feature=gen-gogin type=test control=sequence topic=dos-guard
//ff:what blockBodyLimit — middleware.BodyLimit / MultipartLimit / OverrideBodyLimit 등록
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockBodyLimit_WithMultipartOverride(t *testing.T) {
	doc := buildDoc([]opSpec{{path: "/upload", method: "POST", opID: "Upload"}}, false)
	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{HTTP: &pmanifest.HTTPConfig{
				Overrides: map[string]pmanifest.HTTPOverride{"Upload": {MultipartLimit: "20MiB"}},
			}},
		},
	}
	block := blockBodyLimit(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "multipartOverrides := map[string]int64{") {
		t.Errorf("must emit multipartOverrides map, got:\n%s", body)
	}
	if !strings.Contains(body, `"POST /upload": 20971520,`) {
		t.Errorf("multipart override entry missing, got:\n%s", body)
	}
}
