//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestBindPgtype — UUID/Numeric/Timestamp/Date/Inet/Interval/JSONB/Bytea 의 NestJS 바인딩 매트릭스 테스트

package types

import "testing"

func TestBindPgtype(t *testing.T) {
	runBindMatrix(t, pgtypeBindCases)
}
