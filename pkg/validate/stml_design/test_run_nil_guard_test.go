//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what Run nil guard 테스트 — DesignSpec 또는 STMLPages nil 시 빈 결과

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun_NilDesignSpec(t *testing.T) {
	fs := &yongol.Fullstack{
		STMLPages: []stml.PageSpec{{Name: "page", FileName: "page.html"}},
	}
	if got := Run(fs); len(got) != 0 {
		t.Fatalf("expected 0, got %d", len(got))
	}
}

func TestRun_EmptyPages(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{File: "DESIGN.md"},
	}
	if got := Run(fs); len(got) != 0 {
		t.Fatalf("expected 0, got %d", len(got))
	}
}
