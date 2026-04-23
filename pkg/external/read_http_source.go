//ff:func feature=external type=loader control=sequence
//ff:what readHTTPSource — HTTP(S) URL 을 GET 으로 받아 body 반환 (4xx/5xx 는 body snippet 포함 에러)

package external

import (
	"fmt"
	"io"
	"net/http"
)

// readHTTPSource — GETs an HTTP(S) source, returns body bytes, and includes a body snippet in the error when status >= 400
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
