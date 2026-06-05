//ff:type feature=stml-parse type=model
//ff:what GuardRef — 가드 참조 model.Field (예: workflow.status)
package stml

// GuardRef is a model.Field reference such as "workflow.status".
type GuardRef struct {
	Model string // e.g. "workflow"
	Field string // e.g. "status"
}
