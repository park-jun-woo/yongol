//ff:func feature=external type=util control=sequence
//ff:what HTTP 요청 실행 헬퍼 함수의 코드를 생성한다
package external

import "fmt"

func doHelper(receiverType string) string {
	return fmt.Sprintf(`func (c *%s) do(ctx context.Context, method, path string, body any, result any) error {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return fmt.Errorf("encode request: %%w", err)
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, &buf)
	if err != nil {
		return fmt.Errorf("create request: %%w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("http %%s %%s: %%w", method, path, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return fmt.Errorf("http %%s %%s: status %%d", method, path, resp.StatusCode)
	}
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("decode response: %%w", err)
		}
	}
	return nil
}
`, receiverType)
}
