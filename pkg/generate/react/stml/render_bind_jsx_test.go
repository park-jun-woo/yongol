//ff:func feature=stml-gen type=test control=sequence
//ff:what TestRenderBindJSX — img src 바인딩·boolean/number 타입 분기·미배선 fallback 단위 검증

package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderBindJSX(t *testing.T) {
	ctx := bindCtx{
		all: map[string]map[string]oapiparser.FieldTypeInfo{
			"Op": {
				"can_delete": {Type: "boolean"},
				"thumbnail":  {Type: "string"},
			},
		},
		opID: "Op",
	}

	// <img> binds to src and is self-closing (no text children).
	img := renderBindJSX(stmlparser.FieldBind{Name: "thumbnail", Tag: "img"}, "data", 0, ctx)
	if img != `<img src={data.thumbnail} alt="Thumbnail" />` {
		t.Errorf("img bind: got %q", img)
	}

	// boolean → Yes/No children.
	b := renderBindJSX(stmlparser.FieldBind{Name: "can_delete", Tag: "span"}, "data", 0, ctx)
	if b != "<span>{data.can_delete ? 'Yes' : 'No'}</span>" {
		t.Errorf("boolean bind: got %q", b)
	}

	// Unwired ctx → byte-identical legacy form.
	plain := renderBindJSX(stmlparser.FieldBind{Name: "can_delete", Tag: "span"}, "data", 0, bindCtx{})
	if plain != "<span>{data.can_delete}</span>" {
		t.Errorf("unwired bind: got %q", plain)
	}
}
