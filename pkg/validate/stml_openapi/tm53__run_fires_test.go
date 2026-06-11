//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-53 — object 바인딩 발화 / 스칼라 바인딩 침묵 검증 (배선 확인)

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM53_RunFires(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/thing": getOp("GetThing", nil, map[string]*openapi3.SchemaRef{
			"title": stringProp(),
			"meta":  objectProp(),
		}),
	})

	// object field bound as text → TM-53 fires via Run.
	pages := []stml.PageSpec{{
		Name:     "thing",
		FileName: "thing.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "GetThing",
			Binds:       []stml.FieldBind{{Name: "meta", Tag: "span"}},
		}},
	}}
	if got := countDiag(Run(makeFS(pages, doc)), "[TM-53]"); got != 1 {
		t.Errorf("expected 1 TM-53 via Run, got %d", got)
	}

	// scalar field bound as text → silent.
	pages[0].Fetches[0].Binds = []stml.FieldBind{{Name: "title", Tag: "span"}}
	if got := countDiag(Run(makeFS(pages, doc)), "[TM-53]"); got != 0 {
		t.Errorf("expected 0 TM-53 for a string bind, got %d", got)
	}
}
