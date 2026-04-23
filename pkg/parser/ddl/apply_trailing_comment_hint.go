//ff:func feature=manifest type=parser control=sequence
//ff:what applyTrailingCommentHint — 같은 라인의 `-- @...` 주석을 HintComment 로 파싱해 out 에 추가

package ddl

// applyTrailingCommentHint parses an inline trailing `-- @...` comment and
// appends the resulting hint (if any) to out.
func applyTrailingCommentHint(comment, ddlTrim, path string, lineNum int, tableCtx string, out []HintComment) []HintComment {
	column := extractColumnNameFromLine(ddlTrim)
	hint := parseHintLine("-- "+comment, path, lineNum, tableCtx, column)
	if hint != nil {
		out = append(out, *hint)
	}
	return out
}
