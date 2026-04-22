//ff:func feature=gen-gogin type=util control=sequence
//ff:what logTagForStatus — HTTP status 4xx→"4xx", 그 외→"5xx"

package ssac

// logTagForStatus returns "4xx" / "5xx" used in the slog msg attr.
func logTagForStatus(status int) string {
	if status >= 400 && status < 500 {
		return "4xx"
	}
	return "5xx"
}
