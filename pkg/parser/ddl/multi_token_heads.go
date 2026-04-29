// multiTokenHeads enumerates PostgreSQL multi-word type names that the
// parser must preserve verbatim as the column RawType. The slice is
// written in length-descending order so the greedy match in
// extractRawType picks the longest prefix first (e.g. "TIMESTAMP WITH
// TIME ZONE" wins over a hypothetical shorter "TIMESTAMP WITH" head).
//
// Const-only file — filefunc skips //ff annotations on const/var-only
// files.

package ddl

var multiTokenHeads = [][]string{
	// 4-token forms
	{"TIMESTAMP", "WITH", "TIME", "ZONE"},
	{"TIMESTAMP", "WITHOUT", "TIME", "ZONE"},
	{"TIME", "WITH", "TIME", "ZONE"},
	{"TIME", "WITHOUT", "TIME", "ZONE"},
	// 2-token forms
	{"DOUBLE", "PRECISION"},
	{"CHARACTER", "VARYING"},
	{"BIT", "VARYING"},
}
