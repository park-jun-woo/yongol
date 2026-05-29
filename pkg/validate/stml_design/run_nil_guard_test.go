//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestRun_NilDesignSpec

package stml_design

import (
	"testing"
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
