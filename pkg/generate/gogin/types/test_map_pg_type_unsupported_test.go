//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestMapPGType_Unsupported — 다중 토큰 / CREATE TYPE 거절 회귀

package types

import "testing"

func TestMapPGType_Unsupported(t *testing.T) {
	for _, c := range unsupportedMatrixCases {
		t.Run(c.name, func(t *testing.T) { checkUnsupportedCase(t, c) })
	}
}
