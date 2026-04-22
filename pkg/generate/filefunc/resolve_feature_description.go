//ff:func feature=gen-filefunc type=util control=selection
//ff:what resolveFeatureDescription — picks a feature description: SSOT metadata → infra baseline → fallback
package filefunc

// infraDescriptions holds fixed one-liners for infrastructure packages that
// every Go+Gin backend has, regardless of SSOT content.
var infraDescriptions = map[string]string{
	"api":          "OpenAPI-based Gin router",
	"middleware":   "auth, CORS, and request validation",
	"model":        "DTO structs",
	"auth":         "JWT issuance and verification",
	"db":           "sqlc-generated queries",
	"boot":         "main.go and initialization blocks",
	"service":      "business logic services (SSaC @func implementations)",
	"statemachine": "state transition table",
	"queue":        "queue publish/subscribe adapter",
	"session":      "session backend adapter",
	"cache":        "cache backend adapter",
	"file":         "file storage adapter",
	"authz":        "OPA-based authorization checks",
	"config":       "environment variable and runtime configuration",
	"dashboard":    "aggregation and reporting",
	"report":       "execution reports",
	"schedule":     "cron parser",
	"resolver":     "hierarchical structure resolution",
	"billing":      "credit deduction and lookup",
	"webhook":      "webhook subscription and event publishing",
	"audit":        "audit log",
	"org":          "organization management",
	"user":         "user management",
	"workflow":     "workflow execution, cloning, and run scenarios",
	"template":     "template CRUD and marketplace",
	"execution":    "execution history lookup",
}

// fallbackDescription is used when neither SSOT metadata nor infra baseline
// provides a description for a package.
const fallbackDescription = "backend package"

// resolveFeatureDescription picks the best one-line description for a
// feature using SSOT → infra → fallback priority.
func resolveFeatureDescription(name, ssotDesc string) string {
	switch {
	case ssotDesc != "":
		return ssotDesc
	case hasInfraDescription(name):
		return infraDescriptions[name]
	default:
		return fallbackDescription
	}
}
