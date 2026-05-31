//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"testing"
)

func TestUniqueReverseFieldName_ZeroCov(t *testing.T) {
	if got := uniqueReverseFieldName("posts", 3); got != "posts" {
		t.Errorf("got %q", got)
	}
}
