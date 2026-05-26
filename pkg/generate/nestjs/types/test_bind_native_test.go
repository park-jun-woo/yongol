//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestBindNative — Integer/Float/String/Boolean family 의 NestJS 바인딩 매트릭스 테스트

package types

import "testing"

func TestBindNative(t *testing.T) {
	runBindMatrix(t, nativeBindCases)
}
