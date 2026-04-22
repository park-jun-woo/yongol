//ff:func feature=external type=loader control=sequence
//ff:what readSource — reads OpenAPI source data from a URL or file path
package external

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const httpErrorBodyLimit = 1024 // bytes of response body preserved in error message

func readSource(source string) ([]byte, error) {
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		return readHTTPSource(source)
	}
	return os.ReadFile(source)
}

//ff:func feature=external type=loader control=sequence
//ff:what readHTTPSource — GETs an HTTP(S) source, returns body bytes, and includes a body snippet in the error when status >= 400
func readHTTPSource(source string) ([]byte, error) {
	resp, err := http.Get(source)
	if err != nil {
		return nil, fmt.Errorf("fetch URL %s: %w", source, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		snippet := readBodySnippet(resp.Body, httpErrorBodyLimit)
		if snippet == "" {
			return nil, fmt.Errorf("fetch URL %s: status %d", source, resp.StatusCode)
		}
		return nil, fmt.Errorf("fetch URL %s: status %d: %s", source, resp.StatusCode, snippet)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body from %s: %w", source, err)
	}
	return data, nil
}

//ff:func feature=external type=helper control=sequence
//ff:what readBodySnippet — reads up to limit bytes from a response body and returns a single-line string (best-effort, errors ignored)
func readBodySnippet(r io.Reader, limit int) string {
	buf := make([]byte, limit+1)
	n, _ := io.ReadFull(r, buf)
	if n == 0 {
		return ""
	}
	truncated := n > limit
	if truncated {
		n = limit
	}
	s := strings.ReplaceAll(strings.TrimSpace(string(buf[:n])), "\n", " ")
	if truncated {
		s += "..."
	}
	return s
}
