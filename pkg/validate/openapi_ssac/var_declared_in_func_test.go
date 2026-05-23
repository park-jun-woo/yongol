//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-openapi
//ff:what varDeclaredInFunc — empty varName/미선언/선언됨 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestVarDeclaredInFunc(t *testing.T) {
	fn := ssac.ServiceFunc{
		Sequences: []ssac.Sequence{
			{Result: &ssac.Result{Var: "course"}},
			{Result: nil},
		},
	}

	tests := []struct {
		name    string
		varName string
		want    bool
	}{
		{"empty varName", "", false},
		{"declared var", "course", true},
		{"undeclared var", "user", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := varDeclaredInFunc(fn, tt.varName)
			if got != tt.want {
				t.Errorf("varDeclaredInFunc(%q) = %v, want %v", tt.varName, got, tt.want)
			}
		})
	}
}
