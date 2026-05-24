//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-structural
//ff:what XSS-60 — verifies that @subscribe message field types match the @publish payload types

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xss60SubscribeFieldTypes validates XSS-60: for each topic, the Go type of
// every @publish payload field must be compatible with the corresponding
// @subscribe message struct field type. Mismatches cause a runtime
// json.Unmarshal failure.
func xss60SubscribeFieldTypes(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || len(fs.ServiceFuncs) == 0 {
		return nil
	}

	tableMap := xss60BuildTableMap(fs)

	// Phase 1: build topic → {fieldName → inferredGoType} from @publish.
	publishFieldTypes := xss60CollectPublishTypes(fs, tableMap)
	if len(publishFieldTypes) == 0 {
		return nil
	}

	// Phase 2: compare @subscribe message struct fields against publish types.
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		diags = append(diags, xss60CheckSubscriber(fn, publishFieldTypes)...)
	}
	return diags
}
