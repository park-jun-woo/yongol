//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm32CheckLink — 대상 부재 침묵·구문 단독 보고·필수 충족·매핑 집계 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM32CheckLink(t *testing.T) {
	patterns := map[string]string{"detail": "/d/:AID/:BID?"}
	own := map[string]bool{}

	// Missing target page → silent (TM-31 owns it).
	got := tm32CheckLink(linkRefCtx{Link: &stml.LinkRef{TargetPage: "nope"}}, "f.html", patterns, own)
	if len(got) != 0 {
		t.Errorf("missing target: %+v", got)
	}

	// Syntax error is reported alone.
	got = tm32CheckLink(linkRefCtx{Link: &stml.LinkRef{TargetPage: "detail", ParamsRaw: "garbage"}}, "f.html", patterns, own)
	if len(got) != 1 || !strings.Contains(got[0].Message, "data-link-params") {
		t.Errorf("syntax: %+v", got)
	}

	// Required segment satisfied via mapping, optional unmapped → silent.
	got = tm32CheckLink(linkRefCtx{
		Link:       &stml.LinkRef{TargetPage: "detail", ParamsRaw: "item.id -> AID"},
		InEach:     true,
		ItemFields: map[string]bool{"id": true},
	}, "f.html", patterns, own)
	if len(got) != 0 {
		t.Errorf("satisfied: %+v", got)
	}

	// No params → required :AID unmapped ERROR including the pattern.
	got = tm32CheckLink(linkRefCtx{Link: &stml.LinkRef{TargetPage: "detail"}}, "f.html", patterns, own)
	if len(got) != 1 || !strings.Contains(got[0].Message, "/d/:AID/:BID?") {
		t.Errorf("required unmapped: %+v", got)
	}
}
