//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM10_ComponentRefWithClass_Error

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM10_ComponentRefWithClass_Error(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListItems",
			Components: []stml.ComponentRef{{
				Name:      "DatePicker",
				ClassName: "w-full",
			}},
		}},
	}}
	diags := tm10ClassProhibited(pages)
	if !hasDiag(diags, "[TM-10]") {
		t.Errorf("expected TM-10 for ComponentRef class, got %v", diags)
	}
}
