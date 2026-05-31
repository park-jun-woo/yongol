//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — manifest.* 참조 검증 + XNS-57 memory tx publish 직접 호출
package ssac_manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestByNameXns57MemoryTxPublish_ZeroCov(t *testing.T) {
	// nil fs → nil.
	if d := xns57MemoryTxPublish(nil); d != nil {
		t.Errorf("nil fs should yield nil")
	}
	// backend != memory → nil.
	fsPg := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Queue: &pmanifest.QueueBackend{Backend: "postgres"}}}
	if d := xns57MemoryTxPublish(fsPg); d != nil {
		t.Errorf("postgres backend should yield nil, got %v", d)
	}
	// memory backend + tx-bound publish func → warning.
	fsMem := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{Queue: &pmanifest.QueueBackend{Backend: "memory"}},
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name:     "CreateOrder",
			FileName: "create_order.ssac",
			Sequences: []ssacparser.Sequence{
				{Type: "post", Topic: ""},
				{Type: "publish", Topic: "order.created"},
			},
		}},
	}
	d := xns57MemoryTxPublish(fsMem)
	if len(d) != 1 {
		t.Errorf("memory + tx-bound publish should yield 1 warning, got %d", len(d))
	}
}
