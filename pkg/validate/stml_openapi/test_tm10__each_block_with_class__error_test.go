//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM10_EachBlockWithClass_Error

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM10_EachBlockWithClass_Error(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListItems",
			Eaches: []stml.EachBlock{{
				Field:     "items",
				ClassName: "grid",
			}},
		}},
	}}
	diags := tm10ClassProhibited(pages)
	if !hasDiag(diags, "[TM-10]") {
		t.Errorf("expected TM-10 for EachBlock class, got %v", diags)
	}
}
