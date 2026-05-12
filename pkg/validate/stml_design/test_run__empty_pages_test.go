//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestRun_EmptyPages

package stml_design

import (
	"testing"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun_EmptyPages(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{File: "DESIGN.md"},
	}
	if got := Run(fs); len(got) != 0 {
		t.Fatalf("expected 0, got %d", len(got))
	}
}
