//ff:func feature=chain type=util control=iteration dimension=1
//ff:what toSnakeCase converts a PascalCase string to snake_case.
package chain

func toSnakeCase(s string) string {
	var result []byte
	for i, r := range s {
		if r < 'A' || r > 'Z' {
			result = append(result, byte(r))
			continue
		}
		if i > 0 {
			result = append(result, '_')
		}
		result = append(result, byte(r+'a'-'A'))
	}
	return string(result)
}
