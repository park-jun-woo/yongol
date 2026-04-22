//ff:func feature=manifest type=util control=sequence
//ff:what applyCheckEnum — CHECK enum 값을 Table에 적용
package ddl

func applyCheckEnum(line, colName string, t *Table) {
	col, vals := parseCheckEnum(line)
	if colName != "" {
		col = colName
	}
	if col == "" || len(vals) == 0 {
		return
	}
	if t.CheckEnums == nil {
		t.CheckEnums = map[string][]string{}
	}
	t.CheckEnums[col] = vals
}
