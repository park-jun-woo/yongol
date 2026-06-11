//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TM-53 — 발화(비스칼라·미지원 태그·img 불일치) / 통과(스칼라·img+string·boolean·each 스칼라) 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM53UnrenderableBind(t *testing.T) {
	entry := tm53Entry()

	// --- firing cases: each yields exactly one TM-53 ---
	fire := []struct {
		name string
		bind stml.FieldBind
	}{
		{"object bound as text (a)", stml.FieldBind{Name: "meta", Tag: "span"}},
		{"array bound as text (a)", stml.FieldBind{Name: "tags", Tag: "span"}},
		{"unsupported tag input (b)", stml.FieldBind{Name: "title", Tag: "input"}},
		{"unsupported tag video (b)", stml.FieldBind{Name: "title", Tag: "video"}},
		{"img bound to integer (c)", stml.FieldBind{Name: "count", Tag: "img"}},
	}
	for _, c := range fire {
		f := stml.FetchBlock{OperationID: "GetThing", Binds: []stml.FieldBind{c.bind}}
		if d := tm53UnrenderableBind(f, "GetThing", "p.html", entry, nil); len(d) != 1 {
			t.Errorf("%s: expected 1 TM-53, got %+v", c.name, d)
		}
	}

	// --- passing cases: silent ---
	pass := []struct {
		name string
		bind stml.FieldBind
	}{
		{"string text bind", stml.FieldBind{Name: "title", Tag: "span"}},
		{"img bound to string URL", stml.FieldBind{Name: "avatar", Tag: "img"}},
		{"boolean stays silent", stml.FieldBind{Name: "active", Tag: "span"}},
		{"unknown field owned by TM-06", stml.FieldBind{Name: "ghost", Tag: "input"}},
		{"dotted bind skipped", stml.FieldBind{Name: "meta.x", Tag: "span"}},
	}
	for _, c := range pass {
		f := stml.FetchBlock{OperationID: "GetThing", Binds: []stml.FieldBind{c.bind}}
		if d := tm53UnrenderableBind(f, "GetThing", "p.html", entry, nil); len(d) != 0 {
			t.Errorf("%s: expected no TM-53, got %+v", c.name, d)
		}
	}

	// --- data-each item binds (item schema lookup) ---
	itemTypes := map[string]map[string]map[string]string{
		"GetThing": {"photos": {"url": "string", "size": "integer"}},
	}
	eachCases := []struct {
		name string
		bind stml.FieldBind
		want int
	}{
		{"scalar string item bind silent", stml.FieldBind{Name: "url", Tag: "span"}, 0},
		{"img on integer item fires", stml.FieldBind{Name: "size", Tag: "img"}, 1},
		{"unknown item field silent", stml.FieldBind{Name: "ghost", Tag: "span"}, 0},
	}
	for _, c := range eachCases {
		f := stml.FetchBlock{OperationID: "GetThing", Eaches: []stml.EachBlock{{
			Field: "photos",
			Binds: []stml.FieldBind{c.bind},
		}}}
		if d := tm53UnrenderableBind(f, "GetThing", "p.html", entry, itemTypes); len(d) != c.want {
			t.Errorf("each %s: expected %d TM-53, got %+v", c.name, c.want, d)
		}
	}
}
