//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what baseUrl prefix 와 모든 path 키 공통 접두가 겹치는 이중 접두를 정적 검출한다 (BUG-110)

package react

import (
	"fmt"
	"strings"
)

// detectDoublePrefix reports an error when the non-empty baseURL prefix is
// shared by every path key, which would double-prefix the request URL at
// runtime (baseURL + path = e.g. "/api" + "/api/..." = "/api/api/...").
//
// An empty baseURL or an empty path set is never a double prefix and returns
// nil. The error message includes baseURL, an example path, and the resulting
// doubled URL to make the misconfiguration obvious. See BUG-110.
func detectDoublePrefix(baseURL string, paths []string) error {
	if baseURL == "" || len(paths) == 0 {
		return nil
	}
	for _, path := range paths {
		if !strings.HasPrefix(path, baseURL) {
			return nil
		}
	}
	example := paths[0]
	return fmt.Errorf("double API prefix: baseUrl %q and every path key share prefix %q; request URLs would double-prefix (e.g. %q -> %q)", baseURL, baseURL, example, baseURL+example)
}
