//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"testing"
)

func TestIsPrimaryKey_ZeroCov(t *testing.T) {
	if !isPrimaryKey("id", []string{"id", "tenant_id"}) {
		t.Error("expected id to be primary key")
	}
	if isPrimaryKey("name", []string{"id"}) {
		t.Error("expected name not to be primary key")
	}
}
