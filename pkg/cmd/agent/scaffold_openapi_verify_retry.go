//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what scaffoldOpenAPIVerifyRetry — OpenAPI verify 실패 시 retry 1회차 실행

package agent

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func scaffoldOpenAPIVerifyRetry(yamlDoc *string, offsets *[]pathOffset, pathBlocks map[string]any, pathToOps map[string][]string, opToPath map[string]string, featByOp map[string]features.Feature, newOps map[string]bool, incremental bool, attempt int, cfg Config, out io.Writer, ff *features.FeaturesFile) verifyRetryResult {
	verifyErr := verifyOpenAPI([]byte(*yamlDoc))
	if verifyErr == nil {
		return verifyRetryResult{verified: true}
	}

	allFailed, relativeLines := extractErrorOps(verifyErr, *offsets, ff.Features, *yamlDoc)
	var failedOps []string
	for _, op := range allFailed {
		if !incremental || newOps[op] {
			failedOps = append(failedOps, op)
		}
	}
	if len(failedOps) == 0 {
		fmt.Fprintf(out, "  scaffold openapi: verify error (cannot attribute): %v\n", verifyErr)
		return verifyRetryResult{stopped: true}
	}

	fmt.Fprintf(out, "  scaffold openapi: verify failed (attempt %d/%d), retrying ops: %v\n",
		attempt+1, maxVerifyRetries, failedOps)

	for _, opName := range failedOps {
		retryFailedOp(opName, featByOp, relativeLines, verifyErr, pathBlocks, pathToOps, opToPath, cfg, out)
	}

	*yamlDoc, *offsets = assembleFullOpenAPI("API", pathBlocks, pathToOps)
	return verifyRetryResult{}
}
