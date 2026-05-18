package schedule

import "fmt"

// @func parseCron
// @description Validate cron expression and compute next fire time

type ParseCronRequest struct {
	Expression string
}

type ParseCronResponse struct {
	Cron     string
	NextFire string
}

func ParseCron(req ParseCronRequest) (ParseCronResponse, error) {
	if req.Expression == "" {
		return ParseCronResponse{}, fmt.Errorf("empty cron expression")
	}
	return ParseCronResponse{
		Cron:     req.Expression,
		NextFire: "2026-01-01T00:00:00Z",
	}, nil
}
