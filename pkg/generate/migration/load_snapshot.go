//ff:func feature=migration type=parser control=sequence
//ff:what LoadSnapshot — 스냅샷 파일 읽기 + YONGOL_SCHEMA_HASH 검증 후 AST 반환
package migration

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// LoadSnapshot reads the snapshot file and returns the parsed Schema.
// Returns (nil, nil) when the file does not exist (= initial mode).
// Returns an error when the header hash does not match the body sha256.
func LoadSnapshot(path string) (*Schema, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	nl := strings.Index(text, "\n")
	if nl < 0 {
		return nil, fmt.Errorf("snapshot has no header line")
	}
	head := text[:nl]
	body := text[nl+1:]
	if !strings.HasPrefix(head, SnapshotHashHeaderPrefix) {
		return nil, fmt.Errorf("snapshot header missing %q prefix", SnapshotHashHeaderPrefix)
	}
	stored := strings.TrimSpace(strings.TrimPrefix(head, SnapshotHashHeaderPrefix))
	sum := sha256.Sum256([]byte(body))
	if hex.EncodeToString(sum[:]) != stored {
		return nil, fmt.Errorf("snapshot drift: hash mismatch")
	}
	s := NewSchema()
	if err := BuildASTFromSQL(s, body); err != nil {
		return nil, fmt.Errorf("parse snapshot: %w", err)
	}
	return s, nil
}
