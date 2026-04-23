//ff:func feature=migration type=parser control=iteration dimension=1
//ff:what applyColumnAttrs — 컬럼 뒤 토큰 순회하며 NOT NULL/DEFAULT/PRIMARY KEY/UNIQUE/REFERENCES/CHECK 적용
package migration

// applyColumnAttrs walks the post-type tokens and updates col (and
// table-level constraints on t) accordingly.
func applyColumnAttrs(t *Table, col *Column, rest []string) {
	i := 0
	for i < len(rest) {
		step := dispatchColumnAttr(t, col, rest, i)
		if step <= 0 {
			step = 1
		}
		i += step
	}
}
