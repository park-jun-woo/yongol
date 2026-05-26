//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestBindCompound — Enum/Array/Unsupported 의 NestJS 바인딩 매트릭스 테스트

package types

import "testing"

func TestBindCompound(t *testing.T) {
	runBindMatrix(t, compoundBindCases)
}
