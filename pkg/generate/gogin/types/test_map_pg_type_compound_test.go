//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMapPGType_Compound — JSONB/BYTEA/Array/Enum family 매트릭스 회귀

package types

import "testing"

func TestMapPGType_Compound(t *testing.T) {
	runMatrix(t, compoundMatrixCases)
}
