//ff:func feature=stml-gen type=test control=sequence
//ff:what TestSetLinkTargetsInChildren — ChildNode 슬라이스 순회로 TargetPattern 설정 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSetLinkTargetsInChildren(t *testing.T) {
	link := stmlparser.LinkRef{TargetPage: "building-detail"}
	nodes := []stmlparser.ChildNode{{Kind: "link", Link: &link}}
	setLinkTargetsInChildren(nodes, map[string]string{"building-detail": "/buildings/:BuildingID"})
	if link.TargetPattern != "/buildings/:BuildingID" {
		t.Errorf("TargetPattern = %q", link.TargetPattern)
	}
	// Unknown target resolves to empty (renderer falls back).
	other := stmlparser.LinkRef{TargetPage: "nope"}
	setLinkTargetsInChildren([]stmlparser.ChildNode{{Kind: "link", Link: &other}}, map[string]string{})
	if other.TargetPattern != "" {
		t.Errorf("unknown target: TargetPattern = %q", other.TargetPattern)
	}
}
