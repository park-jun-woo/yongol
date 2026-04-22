//ff:func feature=gen-gogin type=util control=sequence
//ff:what WarnExceeds — 생성 소스의 Q1/Q4 초과 여부를 검사하여 WARN 메시지 리스트 반환

package qcheck

import "fmt"

// WarnExceeds parses src and returns one WARN line per Q1/Q4 excess.
// Parser errors are returned as a single-entry slice with the raw error so
// callers can log the template bug instead of silently swallowing it.
// Depth and loop-body scans are delegated to warnDepthExceeds and
// warnLoopExceeds so each stage stays single-purpose.
func WarnExceeds(filename, src string, lim Limits) []string {
	var warns []string
	depthWarns, derr := warnDepthExceeds(filename, src, lim)
	if derr != nil {
		return []string{fmt.Sprintf("WARN: qcheck parse error in %s: %v", filename, derr)}
	}
	warns = append(warns, depthWarns...)
	loopWarns, lerr := warnLoopExceeds(filename, src, lim)
	if lerr != nil {
		warns = append(warns, fmt.Sprintf("WARN: qcheck loop parse error in %s: %v", filename, lerr))
		return warns
	}
	warns = append(warns, loopWarns...)
	return warns
}
