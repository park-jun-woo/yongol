//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what TestHurlOpenAPIHelpers — unit tests for the pure hurl_openapi helper functions
package hurl_openapi

import (
	"testing"
)

func TestXoh13IsGuardType(t *testing.T) {
	for _, ty := range []string{"empty", "exists", "auth", "state", "eval"} {
		if !xoh13IsGuardType(ty) {
			t.Errorf("%q should be a guard type", ty)
		}
	}
	for _, ty := range []string{"get", "post", "publish", ""} {
		if xoh13IsGuardType(ty) {
			t.Errorf("%q should not be a guard type", ty)
		}
	}
}
