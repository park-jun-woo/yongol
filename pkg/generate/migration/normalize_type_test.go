//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestNormalizeType — aliases·SERIAL·VARCHAR 길이·NUMERIC(p,s)·ARRAY 케이스 전수 점검
package migration

import "testing"

func TestNormalizeType(t *testing.T) {
	for _, c := range normalizeTypeCases() {
		c := c
		t.Run(c.in, func(t *testing.T) {
			assertNormalizeTypeCase(t, c)
		})
	}
}
