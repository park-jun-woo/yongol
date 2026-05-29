//ff:type feature=cli type=test-helper
//ff:what failWriter — 항상 에러를 반환하는 io.Writer (테스트용)

package main

import "fmt"

// failWriter is an io.Writer that always returns an error.
type failWriter struct{}

func (failWriter) Write(p []byte) (int, error) {
	return 0, fmt.Errorf("simulated write failure")
}
