//ff:type feature=validate type=model topic=states
//ff:what statefulTarget — XSM-27 advice 조립용 식별자 묶음

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// statefulTarget collects identifiers XSM-27 needs when composing its advice.
type statefulTarget struct {
	Resource    string // singular lowercase resource name (e.g. "workflow")
	Table       string // plural lowercase DDL table name (e.g. "workflows")
	Diagram     *statemachine.StateDiagram
	StateColumn string // DDL column that holds the state (default "status")
	Model       string // PascalCase model name (e.g. "Workflow")
}
