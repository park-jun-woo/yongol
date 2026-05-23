//ff:func feature=agent type=helper control=sequence
//ff:what opFromSSaCFile — SSaC 파일명에서 operationId 추출

package agent

import (
	"path/filepath"
	"strings"
)

func opFromSSaCFile(relPath string) string {
	base := filepath.Base(relPath)
	return strings.TrimSuffix(base, ".ssac")
}
