//ff:type feature=migration type=test-helper
//ff:what normalizeTypeCase — NormalizeType 케이스 테스트 입력/기대값 타입
package migration

type normalizeTypeCase struct {
	in         string
	wantBase   string
	wantLen    int
	wantPrec   int
	wantScale  int
	wantArray  bool
	wantSerial bool
}
