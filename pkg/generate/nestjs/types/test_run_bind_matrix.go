//ff:func feature=gen-nestjs type=test-helper control=iteration dimension=1
//ff:what runBindMatrix — bindCase 슬라이스를 sub-test 로 실행하는 디스패처

package types

import "testing"

// runBindMatrix iterates the table and delegates each row to
// checkBindCase inside a t.Run sub-test.
func runBindMatrix(t *testing.T, cases []bindCase) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) { checkBindCase(t, c) })
	}
}
