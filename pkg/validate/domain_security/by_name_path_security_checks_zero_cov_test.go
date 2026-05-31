//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestByName_ZeroCov — domain_security 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package domain_security

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestByNamePathSecurityChecks_ZeroCov(t *testing.T) {
	adminItem := &openapi3.PathItem{Get: emptySecurityOp("AdminList")}
	if d := checkAdminPathSecurity("/admin/x", adminItem, "admin.yaml"); len(d) == 0 {
		t.Errorf("checkAdminPathSecurity expected diagnostic for empty security")
	}

	internalItem := &openapi3.PathItem{Post: securedOp("InternalCreate")}
	if d := checkInternalPathSecurity("/internal/x", internalItem, "internal.yaml"); len(d) == 0 {
		t.Errorf("checkInternalPathSecurity expected warning for secured internal op")
	}
}
