//ff:func feature=validate type=util control=iteration dimension=2 topic=ddl-structural
//ff:what allTableContents — CREATE TABLE 블록을 테이블명 → 파일 전체 내용 맵으로 수집
package ddl

// allTableContents merges every CREATE TABLE block keyed by table name across
// all DDL files, so sentinel checks can look up the referenced table's body.
func allTableContents(files []sqlFile) map[string]string {
	out := make(map[string]string)
	for _, f := range files {
		for _, blk := range extractTableBlocks(f) {
			// Append rather than overwrite in case INSERTs sit outside the
			// CREATE TABLE block of the referenced table — include the entire
			// file content so sentinel INSERTs are reachable.
			out[blk.tableName] = f.content
		}
	}
	return out
}
