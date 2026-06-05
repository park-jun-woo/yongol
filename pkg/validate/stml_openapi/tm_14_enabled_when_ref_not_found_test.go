//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TM-14 — data-enabled-when ref의 model 해소: 미해소·정상·빈값 분기 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM14EnabledWhenRefNotFound(t *testing.T) {
	// Page fetch GetWorkflow returns a top-level "workflow" property.
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/workflow": getOp("GetWorkflow", nil, map[string]*openapi3.SchemaRef{
			"workflow": stringProp(),
		}),
	})
	opMap := buildOperationMethodMap(doc)
	fetches := []stml.FetchBlock{{OperationID: "GetWorkflow"}}
	modelFetchMap := buildModelFetchMap(fetches, opMap)

	tests := []struct {
		name      string
		action    stml.ActionBlock
		wantDiags int
	}{
		{
			name:      "model resolves to fetch (ok)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow", EnabledWhen: "workflow.status=draft"},
			wantDiags: 0,
		},
		{
			name:      "model not in any fetch (error)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow", EnabledWhen: "order.status=open"},
			wantDiags: 1,
		},
		{
			name:      "empty enabled-when (skip)",
			action:    stml.ActionBlock{OperationID: "ActivateWorkflow"},
			wantDiags: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := tm14EnabledWhenRefNotFound(tt.action, "p.html", modelFetchMap)
			if got := countDiag(diags, "[TM-14]"); got != tt.wantDiags {
				t.Errorf("TM-14 diags = %d, want %d (%+v)", got, tt.wantDiags, diags)
			}
		})
	}
}
