//ff:func feature=agent type=helper control=sequence
//ff:what tableFromDDLFile — DDL 파일명에서 테이블명 추출

package agent

import (
	"path/filepath"
	"strings"
)

func tableFromDDLFile(relPath string) string {
	base := filepath.Base(relPath)
	return strings.TrimSuffix(base, ".sql")
}
