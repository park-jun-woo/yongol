//ff:func feature=stml-gen type=util control=sequence
//ff:what void 여부와 param 존재에 따라 mutation의 fnParam과 apiArgs를 결정한다
package stml

import (
	"strings"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// resolveMutationArgs determines the fnParam and apiArgs strings for a
// useMutation hook based on whether the operation has a request body (isVoid)
// and path parameters (paramArgs).
func resolveMutationArgs(operationID, paramArgs string, isVoid bool, constraints map[string]map[string]oapiparser.FieldConstraint) (fnParam, apiArgs string) {
	if isVoid {
		fnParam = "()"
		apiArgs = paramArgs // empty when no params
		return fnParam, apiArgs
	}
	fnParam = "(data)"
	apiArgs = "data"
	if paramArgs == "" {
		return fnParam, apiArgs
	}
	inner := strings.TrimPrefix(paramArgs, "{ ")
	inner = strings.TrimSuffix(inner, " }")
	apiArgs = "{ ...data, " + inner + " }"
	// body + path: annotate data with zod schema type when constraints exist
	if fields := lookupConstraints(operationID, constraints); len(fields) > 0 {
		schemaName := toLowerFirst(operationID) + "Schema"
		fnParam = "(data: z.infer<typeof " + schemaName + ">)"
	}
	return fnParam, apiArgs
}
