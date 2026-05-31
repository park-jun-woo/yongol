//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package prepared

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestFileBackendFor_ZeroCov(t *testing.T) {
	mc := &pmanifest.ProjectConfig{File: &pmanifest.FileBackend{Backend: "s3"}}
	if f := fileBackendFor(bnFS(mc, nil)); f == nil || f.Backend != "s3" {
		t.Errorf("declared: %#v", f)
	}
	if f := fileBackendFor(bnFS(nil, []ssac.ServiceFunc{bnCallFunc("file.")})); f == nil || f.Backend != "local" {
		t.Errorf("ssac default: %#v", f)
	}
	if f := fileBackendFor(bnFS(nil, nil)); f != nil {
		t.Errorf("unused should be nil")
	}
}
