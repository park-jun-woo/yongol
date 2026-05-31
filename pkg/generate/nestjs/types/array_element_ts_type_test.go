//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestArrayElementTSType — 배열 element head → TS/Prisma 타입 (4 family + 미지원) 커버
package types

import "testing"

func TestArrayElementTSType_ZeroCov(t *testing.T) {
	for _, h := range []string{"BIGINT", "REAL", "TEXT", "BOOLEAN"} {
		if _, ok := arrayElementTSType(h); !ok {
			t.Errorf("%s should be a supported array element", h)
		}
	}
	if _, ok := arrayElementTSType("UUID"); ok {
		t.Errorf("UUID should not be a supported array element")
	}
}
