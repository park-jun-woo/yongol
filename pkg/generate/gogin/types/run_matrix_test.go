//ff:func feature=gen-gogin type=test-helper control=iteration dimension=1
//ff:what runMatrix — matrixCase 슬라이스를 sub-test 로 실행하는 디스패처

package types

import "testing"

// runMatrix iterates the table and delegates each row to checkMatrixCase
// inside a t.Run sub-test. Extracted so Test* functions in this package
// stay within the F1 single-func and Q4 PURE line budget.
func runMatrix(t *testing.T, cases []matrixCase) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) { checkMatrixCase(t, c) })
	}
}
