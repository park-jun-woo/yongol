//ff:func feature=gen-gogin type=util control=sequence
//ff:what FileNameForMethod — receiver 와 method 이름을 결합해 <receiver>_<method>.go 파일명 생성

package fffile

import "github.com/ettle/strcase"

// FileNameForMethod produces the 1-file-1-func file name for a method by
// combining the receiver type identifier and the method identifier.
//
// When receiverType is empty the method is treated as a plain func and the
// result matches FileNameForFunc. An empty methodName yields "" to signal the
// no-op case.
//
// Examples:
//
//	FileNameForMethod("Server", "ActivateWorkflow") // "server_activate_workflow.go"
//	FileNameForMethod("", "ActivateWorkflow")       // "activate_workflow.go"
func FileNameForMethod(receiverType, methodName string) string {
	if methodName == "" {
		return ""
	}
	if receiverType == "" {
		return strcase.ToSnake(methodName) + ".go"
	}
	return strcase.ToSnake(receiverType) + "_" + strcase.ToSnake(methodName) + ".go"
}
