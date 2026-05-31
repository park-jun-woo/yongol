//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXmv10DeadColorEmpty — xmv10DeadColor 색상 토큰 없을 때 early-return 검증
package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXmv10DeadColorEmpty(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{File: "DESIGN.md", Colors: map[string]string{}},
	}
	if d := xmv10DeadColor(fs, pageTokenRefs{}); d != nil {
		t.Errorf("expected nil when no color tokens, got %+v", d)
	}
}
