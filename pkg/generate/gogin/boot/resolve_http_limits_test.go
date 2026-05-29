//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=dos-guard
//ff:what resolveHTTPLimits — manifest.backend.http 에서 global + per-op limit 추출

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveHTTPLimits_Defaults(t *testing.T) {
	for _, fs := range []*yongol.Fullstack{nil, {}, {Manifest: &pmanifest.ProjectConfig{}}} {
		body, multipart, bo, mo := resolveHTTPLimits(fs)
		if body != defaultBodyLimit || multipart != defaultMultipartLimit {
			t.Errorf("expected defaults, got body=%d multipart=%d", body, multipart)
		}
		if len(bo) != 0 || len(mo) != 0 {
			t.Errorf("expected empty overrides, got bo=%v mo=%v", bo, mo)
		}
	}
}

func TestResolveHTTPLimits_GlobalAndOverride(t *testing.T) {
	doc := buildDoc([]opSpec{{path: "/upload", method: "POST", opID: "Upload"}}, false)
	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{HTTP: &pmanifest.HTTPConfig{
				BodyLimit:      "2MiB",
				MultipartLimit: "64MiB",
				Overrides: map[string]pmanifest.HTTPOverride{
					"Upload": {BodyLimit: "10MiB", MultipartLimit: "100MiB"},
				},
			}},
		},
	}
	body, multipart, bo, mo := resolveHTTPLimits(fs)
	if body != int64(2<<20) {
		t.Errorf("body limit = %d, want %d", body, int64(2<<20))
	}
	if multipart != int64(64<<20) {
		t.Errorf("multipart limit = %d, want %d", multipart, int64(64<<20))
	}
	if bo["POST /upload"] != int64(10<<20) {
		t.Errorf("body override = %d, want %d", bo["POST /upload"], int64(10<<20))
	}
	if mo["POST /upload"] != int64(100<<20) {
		t.Errorf("multipart override = %d, want %d", mo["POST /upload"], int64(100<<20))
	}
}
