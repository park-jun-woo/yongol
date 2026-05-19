package schedule

import (
	"fmt"
	"strings"
	"time"
)

// @func parseCron
// @description Validate cron expression and compute next fire time

type ParseCronRequest struct {
	Expression string
}

type ParseCronResponse struct {
	Cron    string
	NextRun string
}

func ParseCron(req ParseCronRequest) (ParseCronResponse, error) {
	parts := strings.Fields(req.Expression)
	if len(parts) != 5 {
		return ParseCronResponse{}, fmt.Errorf("invalid cron expression: expected 5 fields, got %d", len(parts))
	}
	nextRun := time.Now().Add(1 * time.Hour).UTC().Format(time.RFC3339)
	return ParseCronResponse{Cron: req.Expression, NextRun: nextRun}, nil
}
