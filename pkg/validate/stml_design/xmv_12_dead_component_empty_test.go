//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXmv12DeadComponentEmpty — xmv12DeadComponent 컴포넌트 토큰 없을 때 early-return 검증
package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXmv12DeadComponentEmpty(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{File: "DESIGN.md", Components: map[string]design.ComponentToken{}},
	}
	if d := xmv12DeadComponent(fs, pageTokenRefs{}); d != nil {
		t.Errorf("expected nil when no component tokens, got %+v", d)
	}
}
