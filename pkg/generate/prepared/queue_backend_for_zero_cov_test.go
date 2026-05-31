//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package prepared

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestQueueBackendFor_ZeroCov(t *testing.T) {
	mc := &pmanifest.ProjectConfig{Queue: &pmanifest.QueueBackend{Backend: "postgres"}}
	if q := queueBackendFor(bnFS(mc, nil)); q == nil || q.Backend != "postgres" {
		t.Errorf("declared: %#v", q)
	}
	pub := ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "publish"}}}
	if q := queueBackendFor(bnFS(nil, []ssac.ServiceFunc{pub})); q == nil || q.Backend != "postgres" {
		t.Errorf("ssac default: %#v", q)
	}
	if q := queueBackendFor(bnFS(nil, nil)); q != nil {
		t.Errorf("unused should be nil")
	}
}
