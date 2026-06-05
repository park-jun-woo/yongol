//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestBuildWherePartsEmpty — TestBuildWhereParts — FieldArg 목록 → Prisma where "key: value" 조각 생성 검증

package ssac

import (
	"testing"
)

func TestBuildWherePartsEmpty(t *testing.T) {
	if parts := buildWhereParts(nil); len(parts) != 0 {
		t.Errorf("parts = %v, want empty for no args", parts)
	}
}
