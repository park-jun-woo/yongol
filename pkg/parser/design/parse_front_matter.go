//ff:func feature=frontend type=parser control=sequence
//ff:what --- 구분자 사이의 YAML front matter 와 body 를 분리하여 반환
package design

import (
	"bytes"
	"errors"
)

var errNoFrontMatter = errors.New("DESIGN.md: missing YAML front matter (--- delimiters)")

// parseFrontMatter splits raw file content into YAML front matter and body.
// Front matter must be enclosed by --- lines at the very start of the file.
func parseFrontMatter(data []byte) (yamlPart []byte, body []byte, err error) {
	sep := []byte("---")
	trimmed := data

	// First --- must be the first line.
	if !bytes.HasPrefix(bytes.TrimLeft(trimmed, "\xef\xbb\xbf"), sep) {
		return nil, nil, errNoFrontMatter
	}

	// Find the opening delimiter (skip BOM if present).
	start := bytes.Index(trimmed, sep)
	if start < 0 {
		return nil, nil, errNoFrontMatter
	}
	afterFirst := start + len(sep)

	// Find the closing delimiter.
	rest := trimmed[afterFirst:]
	end := bytes.Index(rest, sep)
	if end < 0 {
		return nil, nil, errNoFrontMatter
	}

	yamlPart = rest[:end]
	body = rest[end+len(sep):]

	// Strip leading newline from body.
	if len(body) > 0 && body[0] == '\n' {
		body = body[1:]
	} else if len(body) > 1 && body[0] == '\r' && body[1] == '\n' {
		body = body[2:]
	}

	return yamlPart, body, nil
}
