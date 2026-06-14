//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what isMutationMethod — POST/PUT/PATCH/DELETE는 mutation, GET 등은 비-mutation 판정 검증
package stml_openapi

import "testing"

func TestIsMutationMethod(t *testing.T) {
	for _, m := range []string{"POST", "PUT", "PATCH", "DELETE"} {
		if !isMutationMethod(m) {
			t.Errorf("%s should be a mutation method", m)
		}
	}
	for _, m := range []string{"GET", "HEAD", "OPTIONS", ""} {
		if isMutationMethod(m) {
			t.Errorf("%s should not be a mutation method", m)
		}
	}
}
