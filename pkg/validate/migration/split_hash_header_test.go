//ff:func feature=validate type=test control=iteration dimension=1 topic=migration-snapshot
//ff:what splitHashHeader — 빈 입력/줄바꿈 없음/prefix 누락/정상 분리 검증

package migration

import (
	"testing"

	gmig "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestSplitHashHeader(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		wantHead string
		wantBody string
		wantOK   bool
	}{
		{
			name:   "empty string returns false",
			text:   "",
			wantOK: false,
		},
		{
			name:   "no newline returns false",
			text:   gmig.SnapshotHashHeaderPrefix + "abc123",
			wantOK: false,
		},
		{
			name:   "no hash prefix returns false",
			text:   "CREATE TABLE t (id INT);\n",
			wantOK: false,
		},
		{
			name:     "valid header splits correctly",
			text:     gmig.SnapshotHashHeaderPrefix + "abc123\nCREATE TABLE t;\n",
			wantHead: gmig.SnapshotHashHeaderPrefix + "abc123",
			wantBody: "CREATE TABLE t;\n",
			wantOK:   true,
		},
		{
			name:     "CRLF normalized to LF",
			text:     gmig.SnapshotHashHeaderPrefix + "abc123\r\nCREATE TABLE t;\r\n",
			wantHead: gmig.SnapshotHashHeaderPrefix + "abc123",
			wantBody: "CREATE TABLE t;\n",
			wantOK:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertSplitHashHeader(t, tt.text, tt.wantOK, tt.wantHead, tt.wantBody)
		})
	}
}
