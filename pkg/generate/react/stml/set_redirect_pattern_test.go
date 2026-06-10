//ff:func feature=stml-gen type=test control=sequence
//ff:what TestSetRedirectPattern — 페이지명 참조만 패턴 해석, 정적 경로·nil 맵은 불변 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSetRedirectPattern(t *testing.T) {
	patterns := map[string]string{"contract-edit": "/contract-edit/:ContractID"}

	// Page-name reference resolves.
	a := stmlparser.ActionBlock{OperationID: "CreateContract", Redirect: "contract-edit"}
	setRedirectPattern(&a, patterns)
	if a.RedirectPattern != "/contract-edit/:ContractID" {
		t.Errorf("page name: RedirectPattern = %q", a.RedirectPattern)
	}

	// Static path stays empty.
	b := stmlparser.ActionBlock{OperationID: "Login", Redirect: "/"}
	setRedirectPattern(&b, patterns)
	if b.RedirectPattern != "" {
		t.Errorf("static: RedirectPattern = %q, want empty", b.RedirectPattern)
	}

	// No redirect stays empty.
	c := stmlparser.ActionBlock{OperationID: "CreateThing"}
	setRedirectPattern(&c, patterns)
	if c.RedirectPattern != "" {
		t.Errorf("no redirect: RedirectPattern = %q, want empty", c.RedirectPattern)
	}

	// Nil map stays empty (renderer falls back to "/<page-name>").
	d := stmlparser.ActionBlock{OperationID: "CreateContract", Redirect: "contract-edit"}
	setRedirectPattern(&d, nil)
	if d.RedirectPattern != "" {
		t.Errorf("nil map: RedirectPattern = %q, want empty", d.RedirectPattern)
	}
}
