//ff:func feature=validate type=util control=sequence topic=migration-snapshot
//ff:what 스냅샷 파일의 YONGOL_SCHEMA_HASH 첫 줄과 본문을 분리

package migration

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// splitHashHeader returns the first line (if it has the hash prefix)
// and the remaining body. ok=false when the file doesn't start with
// the expected prefix.
func splitHashHeader(text string) (string, string, bool) {
	// normalise CRLF to LF so hashes stay stable across Git autocrlf.
	text = strings.ReplaceAll(text, "\r\n", "\n")
	nl := strings.Index(text, "\n")
	if nl < 0 {
		return "", "", false
	}
	head := text[:nl]
	if !strings.HasPrefix(head, migration.SnapshotHashHeaderPrefix) {
		return "", "", false
	}
	return head, text[nl+1:], true
}
