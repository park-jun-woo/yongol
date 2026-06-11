//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseSitemapReader_ReadErrorIsDiag — reader 실패 시 html parse 에러 진단 검증

package stml

import (
	"errors"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestParseSitemapReader_ReadErrorIsDiag(t *testing.T) {
	r := iotest.ErrReader(errors.New("disk gone"))
	spec, diags := ParseSitemapReader("sitemap.html", r)
	if len(diags) != 1 {
		t.Fatalf("expected 1 parse-error diag, got %+v", diags)
	}
	d := diags[0]
	if d.File != "sitemap.html" || d.Phase != diagnostic.PhaseParse || d.Level != diagnostic.LevelError {
		t.Errorf("diag meta = %+v", d)
	}
	if !strings.Contains(d.Message, "html parse:") || !strings.Contains(d.Message, "disk gone") {
		t.Errorf("Message = %q, want html parse: + reader error", d.Message)
	}
	if len(spec.Navs) != 0 || spec.FileName != "" {
		t.Errorf("spec = %+v, want zero value on error", spec)
	}
}
