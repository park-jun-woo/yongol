//ff:func feature=cli type=util control=sequence
//ff:what formatLocation — formats the file:line: prefix, omitting absent parts
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
