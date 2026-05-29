//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestReclassifyCallTargets — pkg.Func 는 pkgCalls, local.Method 는 localMethods 로 분류 검증

package contract

import (
	"reflect"
	"testing"
)

func TestReclassifyCallTargets(t *testing.T) {
	pkgs := map[string]bool{"billing": true}
	calls := []string{
		"billing.CheckCredits", // known pkg → pkgCalls
		"qtx.UserFindByID",     // local recv → localMethods (method name only)
		"noDot",                // dropped
	}
	pkgCalls, localMethods := reclassifyCallTargets(calls, pkgs)
	if !reflect.DeepEqual(pkgCalls, []string{"billing.CheckCredits"}) {
		t.Errorf("pkgCalls = %v", pkgCalls)
	}
	if !reflect.DeepEqual(localMethods, []string{"UserFindByID"}) {
		t.Errorf("localMethods = %v", localMethods)
	}

	t.Run("empty input", func(t *testing.T) {
		p, l := reclassifyCallTargets(nil, pkgs)
		if p != nil || l != nil {
			t.Errorf("expected nil slices, got %v / %v", p, l)
		}
	})
}
