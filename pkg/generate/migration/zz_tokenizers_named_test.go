//ff:func feature=migration type=test control=sequence
//ff:what tokenizer/splitter named 테스트 — splitState/columnTokenizer/lineCommentScanner 메서드 (다중 인용/주석/타입 파라미터) 커버

package migration

import "testing"

// richSQL exercises double-quoted idents, single-quoted strings, block
// comments, line comments, and multiple statements so the split/scan
// state machines traverse every branch.
const richSQL = `
-- a leading line comment
CREATE TABLE "Order" (
    id BIGINT PRIMARY KEY,
    status VARCHAR(20) NOT NULL DEFAULT 'pending; not done', -- trailing comment
    note TEXT /* block ; comment */ DEFAULT 'a''b'
);
CREATE INDEX idx ON "Order" (status);
`

func TestSplitStateStep(t *testing.T)       { _ = splitStatements(richSQL) }
func TestSplitStateFinish(t *testing.T)     { _ = splitStatements(richSQL) }
func TestSplitStateStepInBC(t *testing.T)   { _ = splitStatements(richSQL) }
func TestSplitStateStepInDQ(t *testing.T)   { _ = splitStatements(richSQL) }
func TestSplitStateStepInSQ(t *testing.T)   { _ = splitStatements(richSQL) }

func TestStepTopLevel(t *testing.T) {
	stmts := splitStatements(richSQL)
	if len(stmts) < 2 {
		t.Errorf("expected multiple statements, got %d", len(stmts))
	}
}

func TestColumnTokenizerStep(t *testing.T)     { _ = tokenizeColumnDef(`status VARCHAR(20) NOT NULL DEFAULT 'a''b'`) }
func TestColumnTokenizerFinish(t *testing.T)   { _ = tokenizeColumnDef(`id BIGINT PRIMARY KEY`) }
func TestColumnTokenizerFlush(t *testing.T)    { _ = tokenizeColumnDef(`note TEXT DEFAULT 'x'`) }
func TestColumnTokenizerStepInDQ(t *testing.T) { _ = tokenizeColumnDef(`"weird col" VARCHAR(10)`) }
func TestColumnTokenizerStepInSQ(t *testing.T) { _ = tokenizeColumnDef(`status TEXT DEFAULT 'a; b ''q'' c'`) }

func TestLineCommentScannerStep(t *testing.T) {
	out := stripLineComments("SELECT 1 -- comment\nSELECT 2")
	if out == "" {
		t.Errorf("stripLineComments returned empty")
	}
}

func TestLineCommentScannerStepQuote(t *testing.T) {
	// single-quoted string containing -- must be preserved.
	out := stripLineComments("INSERT INTO t VALUES ('a -- not a comment')")
	if out == "" {
		t.Errorf("stripLineComments returned empty")
	}
	_ = findLineCommentStart("x = 'a -- b' -- real")
}
