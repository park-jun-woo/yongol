//ff:func feature=migration type=test control=sequence
//ff:what TestEmitSnapshotCanonicalIdentity — CanonicalSQL snapshot 에 IDENTITY 절 포함 (round-trip)
package migration

import (
	"strings"
	"testing"
)

func TestEmitSnapshotCanonicalIdentity(t *testing.T) {
	sql := `
CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email TEXT NOT NULL
);
`
	s := NewSchema()
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatalf("parse: %v", err)
	}
	snap := CanonicalSQL(s)
	if !strings.Contains(snap, "id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY") {
		t.Errorf("canonical snapshot should contain IDENTITY clause, got:\n%s", snap)
	}
	// Re-parse the snapshot and ensure IDENTITY survives the round-trip.
	s2 := NewSchema()
	if err := BuildASTFromSQL(s2, snap); err != nil {
		t.Fatalf("round-trip parse: %v", err)
	}
	if s2.Tables["users"].Columns[0].Identity == nil {
		t.Errorf("round-trip lost Identity field")
	}
}
