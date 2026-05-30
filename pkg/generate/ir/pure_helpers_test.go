//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what isCRUDSeq/parseDigits/splitModelMethod/planNeedsTransaction/csrfIsActive 순수 헬퍼

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestIsCRUDSeq(t *testing.T) {
	for _, s := range []string{ssac.SeqGet, ssac.SeqPost, ssac.SeqPut, ssac.SeqDelete} {
		if !isCRUDSeq(s) {
			t.Errorf("%q should be CRUD", s)
		}
	}
	for _, s := range []string{"auth", "call", "response", ""} {
		if isCRUDSeq(s) {
			t.Errorf("%q should not be CRUD", s)
		}
	}
}

func TestParseDigits(t *testing.T) {
	cases := []struct {
		in   string
		want int64
	}{
		{"123", 123},
		{"0", 0},
		{"", -1},
		{"12a", -1},
		{"1 2", -1},
		{"007", 7},
	}
	for _, c := range cases {
		if got := parseDigits(c.in); got != c.want {
			t.Errorf("parseDigits(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

func TestSplitModelMethod(t *testing.T) {
	model, method := splitModelMethod("Course.FindByID")
	if model != "Course" || method != "FindByID" {
		t.Errorf("got (%q,%q)", model, method)
	}
	model, method = splitModelMethod("NoDot")
	if model != "" || method != "NoDot" {
		t.Errorf("no dot got (%q,%q)", model, method)
	}
}

func TestPlanNeedsTransaction(t *testing.T) {
	if !planNeedsTransaction([]Op{{Kind: OpGet}, {Kind: OpPost}}) {
		t.Errorf("post should need transaction")
	}
	if !planNeedsTransaction([]Op{{Kind: OpDelete}}) {
		t.Errorf("delete should need transaction")
	}
	if planNeedsTransaction([]Op{{Kind: OpGet}}) {
		t.Errorf("get-only should not need transaction")
	}
	if planNeedsTransaction(nil) {
		t.Errorf("empty should not need transaction")
	}
}

func TestCsrfIsActive(t *testing.T) {
	// csrf not required -> false
	if csrfIsActive(&prepared.State{Auth: prepared.Auth{CsrfRequired: false}}) {
		t.Errorf("not required should be false")
	}
	// required but auth absent -> false
	if csrfIsActive(&prepared.State{Auth: prepared.Auth{CsrfRequired: true, Present: false}}) {
		t.Errorf("auth absent should be false")
	}
	// required, present, raw csrf nil -> default true
	ps := &prepared.State{Auth: prepared.Auth{CsrfRequired: true, Present: true, Raw: &manifest.Auth{}}}
	if !csrfIsActive(ps) {
		t.Errorf("nil csrf with cookie/hybrid should default true")
	}
	// required, present, explicit csrf.Enabled=false -> false
	ps2 := &prepared.State{Auth: prepared.Auth{CsrfRequired: true, Present: true,
		Raw: &manifest.Auth{Csrf: &manifest.CsrfConfig{Enabled: false}}}}
	if csrfIsActive(ps2) {
		t.Errorf("explicit disabled csrf should be false")
	}
}
