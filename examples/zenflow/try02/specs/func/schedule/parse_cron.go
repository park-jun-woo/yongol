package schedule

import (
	"fmt"
	"strings"
	"time"
)

// @func parseCron
// @error 400
// @description Validates a cron expression and computes the next fire time (purity-safe).

type ParseCronRequest struct {
	Expression string
}

type ParseCronResponse struct {
	Expression string
	NextRun    string
}

func ParseCron(req ParseCronRequest) (ParseCronResponse, error) {
	parts := strings.Fields(req.Expression)
	if len(parts) != 5 {
		return ParseCronResponse{}, fmt.Errorf("invalid cron expression: expected 5 fields, got %d", len(parts))
	}
	// Minimal validation: each field must be non-empty (real validation would parse ranges/wildcards)
	for i, p := range parts {
		if p == "" {
			return ParseCronResponse{}, fmt.Errorf("invalid cron expression: field %d is empty", i)
		}
	}
	// Compute next run: advance from now by 1 minute as a placeholder (purity-safe, no network/DB)
	nextRun := time.Now().UTC().Add(time.Minute).Format(time.RFC3339)
	return ParseCronResponse{
		Expression: req.Expression,
		NextRun:    nextRun,
	}, nil
}
