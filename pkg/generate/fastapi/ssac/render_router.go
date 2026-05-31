//ff:func feature=gen-fastapi type=generator control=sequence
//ff:what RenderRouter — feature 단위 FastAPI Router Python 소스 생성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// RenderRouter produces a FastAPI router file for a given feature. Each
// ServicePlan contributes one route handler decorated with the appropriate
// HTTP method decorator. Parameters are typed using PathParams, QueryParams,
// and BodyFields from the ServicePlan.
func RenderRouter(feature string, plans []*ir.ServicePlan) (string, error) {
	if feature == "" {
		return "", fmt.Errorf("RenderRouter: empty feature name")
	}

	var b strings.Builder
	needsAuth, needsEventBus := routerDependencyFlags(plans)
	writeRouterImports(&b, feature, plans, needsAuth, needsEventBus)
	b.WriteString(fmt.Sprintf("router = APIRouter(prefix=\"/%s\", tags=[\"%s\"])\n\n", feature, feature))
	writeRouterHandlers(&b, plans)

	return b.String(), nil
}
