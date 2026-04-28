//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what joinTypes — comma-joins return-type strings for diagnostic messages

package ssac

// joinTypes formats a return-type slice (e.g. []string{"FooResponse",
// "error"}) as "FooResponse, error" for embedding in a diagnostic message.
func joinTypes(types []string) string {
	out := ""
	for i, t := range types {
		if i > 0 {
			out += ", "
		}
		out += t
	}
	return out
}
