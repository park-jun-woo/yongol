//ff:func feature=validate type=util control=iteration dimension=1 topic=func-check
//ff:what joinReturnTypes — comma-joins return-type strings for diagnostic messages

package ssac_func

func joinReturnTypes(types []string) string {
	out := ""
	for i, t := range types {
		if i > 0 {
			out += ", "
		}
		out += t
	}
	return out
}
