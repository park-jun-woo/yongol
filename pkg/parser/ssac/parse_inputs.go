//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what parseInputs — parses {key: value, ...} form input and returns a map
package ssac

import (
	"fmt"
	"strings"
)

// parseInputs parses the {key: value, ...} form.
func parseInputs(s string) (map[string]string, error) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "{")
	s = strings.TrimSuffix(s, "}")
	s = strings.TrimSpace(s)
	if s == "" {
		return map[string]string{}, nil
	}
	result := make(map[string]string)
	for _, pair := range strings.Split(s, ",") {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}
		colonIdx := strings.IndexByte(pair, ':')
		if colonIdx < 0 {
			return nil, fmt.Errorf("%q is not a valid input pair; use \"{Key: value}\" format", pair)
		}
		key := strings.TrimSpace(pair[:colonIdx])
		val := strings.TrimSpace(pair[colonIdx+1:])
		if key != "" && val != "" {
			result[key] = val
		}
	}
	return result, nil
}
