//ff:func feature=gen-ir type=util control=sequence
//ff:what splitModelMethod -- "Model.Method" 문자열을 모델명과 메서드명으로 분리

package ir

import "strings"

// splitModelMethod splits "Course.FindByID" into ("Course", "FindByID").
// If there is no dot, model is empty and method is the full string.
func splitModelMethod(modelMethod string) (model, method string) {
	if idx := strings.IndexByte(modelMethod, '.'); idx >= 0 {
		return modelMethod[:idx], modelMethod[idx+1:]
	}
	return "", modelMethod
}
