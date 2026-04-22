//ff:func feature=validate type=util control=iteration dimension=1 topic=funcspec-structural
//ff:what isForbiddenImport — 주어진 import 경로가 XFF-41 금지 목록에 해당하는지

package funcspec

import "strings"

// forbiddenImportPrefixes: func/ is pure business logic; I/O must be in model/queue/etc.
// io, bufio, os are allowed (bytes/string/tempfile processing).
var forbiddenImportPrefixes = []string{
	"database/sql",
	"net/http",
	"net/rpc",
	"google.golang.org/grpc",
}

func isForbiddenImport(imp string) bool {
	for _, p := range forbiddenImportPrefixes {
		if imp == p || strings.HasPrefix(imp, p+"/") {
			return true
		}
	}
	return false
}
