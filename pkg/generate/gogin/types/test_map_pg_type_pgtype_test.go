//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMapPGType_Pgtype — UUID/NUMERIC/TIMESTAMP/INET/INTERVAL family 매트릭스 회귀

package types

import "testing"

func TestMapPGType_Pgtype(t *testing.T) {
	runMatrix(t, pgtypeMatrixCases)
}
