//ff:func feature=cli type=util control=sequence
//ff:what formatLocation — file:line: prefix 포맷 (file/line 부재 시 생략)
package main

import "fmt"

// formatLocation returns a "file:line: " prefix for diagnostic output.
// Empty string when file is unknown. Line is omitted when 0.
func formatLocation(file string, line int) string {
	if file == "" {
		return ""
	}
	if line > 0 {
		return fmt.Sprintf("%s:%d: ", file, line)
	}
	return fmt.Sprintf("%s: ", file)
}
