//ff:func feature=validate type=rule control=sequence topic=states
//ff:what writeXsm27OptionB — XSM-27 advice Option B (state-neutral) 라인 작성

package ssac_statemachine

import (
	"strings"
)

// writeXsm27OptionB writes the "Option B (state-neutral)" stanza that
// tells the author to add `// @state-neutral` when the operation truly
// does not depend on the resource's state.
func writeXsm27OptionB(b *strings.Builder) {
	b.WriteString("Option B (state-neutral): if this operation truly does not depend on the resource's state, add above the function:\n")
	b.WriteString("    // @state-neutral")
}
