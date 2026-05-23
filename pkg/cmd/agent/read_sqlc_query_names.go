//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what readSQLcQueryNames — sqlc 쿼리 파일에서 -- name: 선언 목록 추출

package agent

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// readSQLcQueryNames reads the sqlc query file for a table and returns
// all "-- name:" lines (query name declarations).
func readSQLcQueryNames(specsDir, tableName string) []string {
	if tableName == "" {
		return nil
	}
	path := filepath.Join(specsDir, "db", "queries", tableName+".sql")
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	var names []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "-- name:") {
			names = append(names, line)
		}
	}
	return names
}
