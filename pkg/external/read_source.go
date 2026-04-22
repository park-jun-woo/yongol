//ff:func feature=external type=loader control=sequence
//ff:what URL 또는 파일 경로로부터 OpenAPI 소스 데이터를 읽어온다
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
//ff:what HTTP(S) 소스를 GET 하여 본문 바이트를 반환하고, status>=400 이면 body 스니펫을 에러 메시지에 포함한다
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
//ff:what 응답 body 를 최대 limit 바이트까지 읽어 한 줄 문자열로 반환한다(best-effort, 에러 무시)
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
