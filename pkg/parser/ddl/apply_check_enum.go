//ff:func feature=manifest type=util control=sequence
//ff:what applyCheckEnum — CHECK enum 값을 Column에 적용
package ddl

// applyCheckEnum captures CHECK (col IN (...)) values directly on the
// owning Column (inline case). The current column is the authoritative
// column name regardless of what parseCheckEnum's regex picked up.
func applyCheckEnum(line string, col *Column) {
	_, vals := parseCheckEnum(line)
	if len(vals) == 0 {
		return
	}
	col.CheckEnum = vals
}
