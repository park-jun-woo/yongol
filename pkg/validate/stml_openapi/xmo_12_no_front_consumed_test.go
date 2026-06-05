//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO12NoFrontConsumed — frontend-off/nil-doc 스킵, no-front 태그인데 STML 소비 시 XMO-12 WARNING 생성 검증
package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXMO12NoFrontConsumed(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "users.html",
		Fetches:  []stml.FetchBlock{{OperationID: "ListUsers"}},
	}}
	noFrontOp := &openapi3.PathItem{Get: &openapi3.Operation{OperationID: "ListUsers", Tags: []string{noFrontTag}}}
	doc := makeDoc(map[string]*openapi3.PathItem{"/users": noFrontOp})

	// Frontend ON + consumed no-front op → one XMO-12 WARNING.
	diags := xmo12NoFrontConsumed(makeFS(pages, doc))
	if countDiag(diags, "[XMO-12]") != 1 {
		t.Fatalf("expected 1 XMO-12, got %d (%v)", countDiag(diags, "[XMO-12]"), diags)
	}
	if diags[0].Level != diagnostic.LevelWarning {
		t.Errorf("level = %q, want WARNING", diags[0].Level)
	}
	if !strings.Contains(diags[0].Message, "ListUsers") {
		t.Errorf("message should name ListUsers, got %q", diags[0].Message)
	}

	// Frontend OFF → skipped (nil).
	off := makeFS(pages, doc)
	off.Manifest = &manifest.ProjectConfig{}
	if got := xmo12NoFrontConsumed(off); got != nil {
		t.Errorf("frontend off: expected nil, got %v", got)
	}

	// Frontend ON but nil OpenAPI doc → skipped (nil).
	noDoc := &yongol.Fullstack{
		STMLPages: pages,
		Manifest:  &manifest.ProjectConfig{Frontend: manifest.Frontend{Lang: "typescript"}},
	}
	if got := xmo12NoFrontConsumed(noDoc); got != nil {
		t.Errorf("nil doc: expected nil, got %v", got)
	}
}
