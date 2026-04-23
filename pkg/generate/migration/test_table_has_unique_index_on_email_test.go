//ff:func feature=migration type=test-helper control=iteration dimension=1
//ff:what tableHasUniqueIndexOnEmail — email 단일 컬럼 UNIQUE 인덱스 존재 여부
package migration

func tableHasUniqueIndexOnEmail(tbl *Table) bool {
	for _, idx := range tbl.Indexes {
		if idx.Unique && len(idx.Columns) == 1 && idx.Columns[0] == "email" {
			return true
		}
	}
	return false
}
