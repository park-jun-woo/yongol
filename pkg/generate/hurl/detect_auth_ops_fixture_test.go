//ff:type feature=gen-hurl type=test-helper
//ff:what detectAuthOpsFixture — TestDetectAuthOps 테이블 로우 타입

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// detectAuthOpsFixture is the table-row shape used by TestDetectAuthOps.
// Kept unexported + package-local so sibling test files (new_password_op,
// ssac fixture builders) remain untouched.
type detectAuthOpsFixture struct {
	name        string
	opID        string
	public      bool
	hasPassword bool
	funcBuilder func(opID string) *ssac.ServiceFunc
	wantRole    string // "signup" | "login" | "" (neither)
	wantWarnSub string // substring expected in warnings; "" = no specific warning
}
