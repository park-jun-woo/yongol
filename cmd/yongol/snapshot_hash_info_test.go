//ff:func feature=cli type=test control=sequence
//ff:what TestSnapshotHashInfo — snapshotHashInfo 해시 일치/불일치/파싱 에러 검증

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestSnapshotHashInfo(t *testing.T) {
	t.Run("ValidNoDrift", func(t *testing.T) {
		body := "CREATE TABLE foo ();\n"
		sum := sha256.Sum256([]byte(body))
		hash := hex.EncodeToString(sum[:])
		data := []byte(migration.SnapshotHashHeaderPrefix + hash + "\n" + body)
		stored, drift := snapshotHashInfo(data)
		if drift {
			t.Error("expected no drift")
		}
		if stored != hash {
			t.Errorf("stored = %q, want %q", stored, hash)
		}
	})

	t.Run("Drift", func(t *testing.T) {
		data := []byte(migration.SnapshotHashHeaderPrefix + "badhash\nsome body\n")
		_, drift := snapshotHashInfo(data)
		if !drift {
			t.Error("expected drift for mismatched hash")
		}
	})

	t.Run("NoNewline", func(t *testing.T) {
		data := []byte("single line no newline")
		_, drift := snapshotHashInfo(data)
		if !drift {
			t.Error("expected drift for no newline")
		}
	})

	t.Run("WrongPrefix", func(t *testing.T) {
		data := []byte("-- WRONG PREFIX\nbody\n")
		_, drift := snapshotHashInfo(data)
		if !drift {
			t.Error("expected drift for wrong prefix")
		}
	})
}
