//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what selectColsContain — snake_case 컬럼명이 SelectCols 목록에 존재하는지 확인

package ssac_sqlc

// selectColsContain checks if a snake_case column name exists in the SelectCols list.
func selectColsContain(cols []string, target string) bool {
	for _, c := range cols {
		if c == target {
			return true
		}
	}
	return false
}
