//ff:func feature=ssac-parse type=parser control=sequence
//ff:what parseLine — parses a single @annotation line and returns a Sequence
package ssac

import "strings"

// parseLine parses a single line and returns a Sequence.
// For @response {, returns (nil, true, nil) to signal the start of multi-line mode.
func parseLine(line string) (*Sequence, bool, error) {
	if strings.HasPrefix(line, "@response") {
		return parseResponseLine(line)
	}

	// @type! — detect the ! suffix
	suppressWarn := false
	if idx := strings.IndexByte(line, '!'); idx > 0 {
		spaceIdx := strings.IndexByte(line, ' ')
		if spaceIdx < 0 || idx < spaceIdx {
			line = line[:idx] + line[idx+1:]
			suppressWarn = true
		}
	}

	seq, err := parseAnnotation(line)
	if err != nil {
		return nil, false, err
	}
	if seq != nil && suppressWarn {
		seq.SuppressWarn = true
	}
	return seq, false, nil
}
