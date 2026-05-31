//ff:func feature=validate type=test control=iteration dimension=1 topic=states
//ff:what TestFindByIDResultVar — findByIDResultVar 분기별 결과 변수 탐색 검증
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestFindByIDResultVar(t *testing.T) {
	cases := []struct {
		name    string
		seqs    []ssac.Sequence
		model   string
		wantVar string
		wantOK  bool
	}{
		{
			name:   "empty model returns false",
			seqs:   nil,
			model:  "",
			wantOK: false,
		},
		{
			name:   "wrong type skipped, not found",
			seqs:   []ssac.Sequence{{Type: "post", Model: "Order.FindByID"}},
			model:  "Order",
			wantOK: false,
		},
		{
			name:   "model mismatch skipped, not found",
			seqs:   []ssac.Sequence{{Type: "get", Model: "Other.FindByID"}},
			model:  "Order",
			wantOK: false,
		},
		{
			name:    "match with result var",
			seqs:    []ssac.Sequence{{Type: "get", Model: "Order.FindByID", Result: &ssac.Result{Var: "order"}}},
			model:   "Order",
			wantVar: "order",
			wantOK:  true,
		},
		{
			name:    "match without result returns empty ok",
			seqs:    []ssac.Sequence{{Type: "get", Model: "Order.FindByID"}},
			model:   "Order",
			wantVar: "",
			wantOK:  true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotVar, gotOK := findByIDResultVar(tc.seqs, tc.model)
			if gotVar != tc.wantVar || gotOK != tc.wantOK {
				t.Errorf("findByIDResultVar() = (%q, %v), want (%q, %v)", gotVar, gotOK, tc.wantVar, tc.wantOK)
			}
		})
	}
}
