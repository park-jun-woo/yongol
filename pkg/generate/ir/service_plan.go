//ff:type feature=gen-ir type=model
//ff:what ServicePlan -- 하나의 SSaC 함수에 대한 프레임워크 비의존 중간 표현

package ir

// ServicePlan is the framework-agnostic intermediate representation for a
// single SSaC service function (one endpoint or subscribe handler). Backend
// code generators (gogin, nestjs, fastapi) consume this plan to produce
// framework-specific source code.
type ServicePlan struct {
	// OperationID is the unique operation identifier matching the OpenAPI
	// operationId.
	OperationID string

	// TriggerKind indicates whether this function is triggered by HTTP or a
	// queue subscription.
	TriggerKind TriggerKind

	// HTTPMethod is the HTTP method (e.g. "GET", "POST"). Only used when
	// TriggerKind == TriggerHTTP.
	HTTPMethod string

	// URLPath is the URL path pattern (e.g. "/courses/:id"). Only used when
	// TriggerKind == TriggerHTTP.
	URLPath string

	// Topic is the queue topic (e.g. "order.completed"). Only used when
	// TriggerKind == TriggerSubscribe.
	Topic string

	// FileName is the source SSaC file path (for diagnostics).
	FileName string

	// Feature is the feature folder name (e.g. "auth") that determines the
	// output directory.
	Feature string

	// Imports lists external packages referenced by @call sequences.
	Imports []string

	// Ops is the ordered list of abstract operations to execute.
	Ops []Op

	// UsesTransaction is true when the service function contains at least one
	// mutating operation (@post/@put/@delete) and therefore requires a database
	// transaction.
	UsesTransaction bool

	// QueryMethods lists all database query methods used by this plan.
	QueryMethods []QueryMethod
}
