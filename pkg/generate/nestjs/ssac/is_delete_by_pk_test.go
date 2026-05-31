//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestIsDeleteByPK(t *testing.T) {
	if !isDeleteByPK(nil) {
		t.Errorf("empty args should be PK")
	}
	if !isDeleteByPK([]ir.FieldArg{{Key: "id", IsPK: true}}) {
		t.Errorf("IsPK flag should be PK")
	}
	if !isDeleteByPK([]ir.FieldArg{{Key: "id"}}) {
		t.Errorf("key=id heuristic should be PK")
	}
	if isDeleteByPK([]ir.FieldArg{{Key: "slug"}}) {
		t.Errorf("non-pk key should not be PK")
	}
}
