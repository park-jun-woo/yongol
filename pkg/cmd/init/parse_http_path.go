//ff:func feature=cli-init type=util control=selection
//ff:what parseHTTPPath — splits "POST /workflows/{id}" into method and URI

package cliinit

import (
	"fmt"
	"strings"
)

// parseHTTPPath splits a path like "POST /workflows/{id}" into method and URI.
func parseHTTPPath(path string) (httpRoute, error) {
	parts := strings.Fields(path)
	if len(parts) != 2 {
		return httpRoute{}, fmt.Errorf("invalid path %q: expected 'METHOD /uri'", path)
	}
	method := strings.ToLower(parts[0])
	switch method {
	case "get", "post", "put", "patch", "delete":
		// valid
	default:
		return httpRoute{}, fmt.Errorf("invalid HTTP method %q in path %q", parts[0], path)
	}
	return httpRoute{Method: method, URI: parts[1]}, nil
}
