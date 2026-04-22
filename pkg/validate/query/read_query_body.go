//ff:func feature=validate type=util control=iteration dimension=1 topic=query-structural
//ff:what readQueryBody — QuerySpec에 대응하는 SQL 본문 + 헤더 라인을 파일에서 읽는다

package query

import (
	"bufio"
	"os"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

// readQueryBody reads the SQL body of one query from its source file.
// startLine is QuerySpec.Line (1-based, pointing at the `-- name:` comment).
func readQueryBody(q sqlc.QuerySpec) (*queryBody, error) {
	f, err := os.Open(q.File)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	body := &queryBody{Escapes: make(map[string]bool)}
	scanner := bufio.NewScanner(f)
	lineNo := 0
	inQuery := false
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		if done := readQueryBodyStep(body, line, lineNo, q.Line, &inQuery); done {
			break
		}
	}
	body.Text = strings.Join(body.Lines, "\n")
	return body, scanner.Err()
}
