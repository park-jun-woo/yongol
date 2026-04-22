//ff:func feature=validate type=util control=iteration dimension=1 topic=ddl-structural
//ff:what scanPositionals — 쿼리 파일에서 $N 위치 파라미터 라인/토큰 수집

package ddl

import (
	"bufio"
	"os"
	"regexp"
)

// scanPositionals scans the query body starting at startLine in filePath and
// returns every `$N` positional parameter occurrence until the next
// `-- name:` header or EOF.
func scanPositionals(filePath string, startLine int) []positionalHit {
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer f.Close()

	nameRe := regexp.MustCompile(`^\s*--\s*name:`)
	scanner := bufio.NewScanner(f)
	lineNo := 0
	inQuery := false
	var hits []positionalHit

	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		if lineNo == startLine {
			inQuery = true
			continue
		}
		if !inQuery {
			continue
		}
		if nameRe.MatchString(line) {
			break
		}
		if m := positionalParamRe.FindString(line); m != "" {
			hits = append(hits, positionalHit{line: lineNo, param: m})
		}
	}
	return hits
}
