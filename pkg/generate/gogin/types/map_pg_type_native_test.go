//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMapPGType_Native — Integer/Float/String/Boolean family 매트릭스 회귀

package types

import "testing"

func TestMapPGType_Native(t *testing.T) {
	runMatrix(t, nativeMatrixCases)
}
