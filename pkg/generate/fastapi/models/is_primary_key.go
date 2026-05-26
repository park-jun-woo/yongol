//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what isPrimaryKey — 컬럼이 프라이머리 키에 포함되는지 확인

package models

// isPrimaryKey checks if the column is in the primary key list.
func isPrimaryKey(colName string, pk []string) bool {
	for _, k := range pk {
		if k == colName {
			return true
		}
	}
	return false
}
