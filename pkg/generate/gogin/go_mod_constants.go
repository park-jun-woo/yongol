package gogin

// goModTidyStderrLimit caps captured stderr bytes included in the error
// message. `go mod tidy` and friends can emit multi-KB diagnostics;
// truncating keeps the error readable while still surfacing the first
// concrete cause.
const goModTidyStderrLimit = 4 * 1024

// OTel dependency versions pinned centrally so otel-core, otelgin, otelsql,
// and the exporters stay on a compatible set. Bumps go through a single
// sweep rather than scattered literals across multiple generator blocks.
// Only this co-release family is pinned — everything else uses @latest and
// is frozen per-project in go.sum after the first generate.
const (
	otelVersion       = "v1.32.0"
	otelContribGinVer = "v0.57.0"
	otelSQLVer        = "v0.37.0"
)

// coreDeps are the non-OTel runtime dependencies of the generated backend.
// They use @latest at `go get` time — semver-major pins live in the import
// path (e.g. `jwt/v5`), so @latest only picks minor/patch upgrades. The
// generated project's own go.sum freezes the exact resolution per build.
var coreDeps = []string{
	"github.com/gin-gonic/gin@latest",
	"github.com/gin-contrib/cors@latest",
	"github.com/lib/pq@latest",
	"github.com/golang-jwt/jwt/v5@latest",
	"github.com/oapi-codegen/runtime@latest",
	"github.com/getkin/kin-openapi@latest",
	"github.com/ulule/limiter/v3@latest",
	"github.com/prometheus/client_golang@latest",
	"github.com/oklog/ulid/v2@latest",
	"github.com/park-jun-woo/ssac@latest",
}
