//ff:func feature=cli type=util control=sequence
//ff:what snapshotHashInfo — migration snapshot 헤더에서 stored hash + drift 여부 반환
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// snapshotHashInfo returns (stored hash, drift-boolean) for a snapshot
// file. A parse error is reported as drift to surface loudly.
func snapshotHashInfo(data []byte) (string, bool) {
	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	nl := strings.Index(text, "\n")
	if nl < 0 {
		return "", true
	}
	head := text[:nl]
	body := text[nl+1:]
	if !strings.HasPrefix(head, migration.SnapshotHashHeaderPrefix) {
		return "", true
	}
	stored := strings.TrimSpace(strings.TrimPrefix(head, migration.SnapshotHashHeaderPrefix))
	sum := sha256.Sum256([]byte(body))
	return stored, hex.EncodeToString(sum[:]) != stored
}
