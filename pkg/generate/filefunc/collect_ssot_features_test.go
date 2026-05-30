//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestCollectSSOTFeatures — SSaC/funcspec feature 수집 + nil 처리 검증

package filefunc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCollectSSOTFeatures(t *testing.T) {
	t.Run("NilReturnsEmpty", func(t *testing.T) {
		if got := collectSSOTFeatures(nil); len(got) != 0 {
			t.Errorf("expected empty map for nil fs, got: %v", got)
		}
	})

	t.Run("MergesServiceAndFuncSpec", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Feature: "workflow"},
			},
			ProjectFuncSpecs: []funcspec.FuncSpec{
				{Package: "auth", Description: "authentication"},
			},
		}
		got := collectSSOTFeatures(fs)
		if _, ok := got["workflow"]; !ok {
			t.Errorf("expected workflow from ServiceFuncs: %v", got)
		}
		if got["auth"] != "authentication" {
			t.Errorf("expected auth description, got %q", got["auth"])
		}
	})
}
