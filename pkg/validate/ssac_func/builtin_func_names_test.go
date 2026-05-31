//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestSSaCFuncHelpers — unit tests for the pure ssac_func helper functions
package ssac_func

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestBuiltinFuncNames(t *testing.T) {
	specs := []funcspec.FuncSpec{
		{Package: "auth", Name: "hashPassword"},
		{Package: "auth", Name: "verifyPassword"},
		{Package: "cache", Name: "get"},
	}
	got := builtinFuncNames("auth", specs)
	want := []string{"HashPassword", "VerifyPassword"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("builtinFuncNames(auth) = %v, want %v", got, want)
	}
	if got := builtinFuncNames("missing", specs); got != nil {
		t.Errorf("missing pkg → %v, want nil", got)
	}
}
