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

func TestBlockBodyLimit_WithOverride(t *testing.T) {
	doc := buildDoc([]opSpec{{path: "/upload", method: "POST", opID: "Upload"}}, false)
	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{HTTP: &pmanifest.HTTPConfig{
				Overrides: map[string]pmanifest.HTTPOverride{"Upload": {BodyLimit: "10MiB"}},
			}},
		},
	}
	block := blockBodyLimit(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "r.Use(middleware.OverrideBodyLimit(bodyOverrides, multipartOverrides))") {
		t.Errorf("must emit OverrideBodyLimit when overrides present, got:\n%s", body)
	}
	if !strings.Contains(body, `"POST /upload": 10485760,`) {
		t.Errorf("override entry missing, got:\n%s", body)
	}
}

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
