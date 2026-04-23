//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what parseHintLine — `-- @<tag> key=val ...` 라인을 파싱해 HintComment 반환

package ddl

import "strings"

// parseHintLine accepts a line guaranteed to start with `--`.
// Returns a HintComment if the comment starts with `@<tag>`.
func parseHintLine(line, file string, lineNum int, tableCtx, columnCtx string) *HintComment {
	body := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "--"))
	if !strings.HasPrefix(body, "@") {
		return nil
	}
	body = strings.TrimPrefix(body, "@")
	toks := strings.Fields(body)
	if len(toks) == 0 {
		return nil
	}
	tag := strings.ToLower(toks[0])
	args := map[string]string{}
	for _, t := range toks[1:] {
		if eq := strings.Index(t, "="); eq > 0 {
			args[strings.ToLower(t[:eq])] = t[eq+1:]
		}
	}
	return &HintComment{
		File:      file,
		Line:      lineNum,
		Tag:       tag,
		Args:      args,
		TableCtx:  tableCtx,
		ColumnCtx: columnCtx,
	}
}
