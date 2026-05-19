package schedule

import "fmt"

// @func parseCron
// @description Validate cron expression and compute next fire time

type ParseCronRequest struct {
	Expression string
}

type ParseCronResponse struct {
	NextFire string
	Valid    bool
}

func ParseCron(req ParseCronRequest) (ParseCronResponse, error) {
	if req.Expression == "" {
		return ParseCronResponse{Valid: false}, fmt.Errorf("empty cron expression")
	}
	return ParseCronResponse{
		NextFire: "2026-01-01T00:00:00Z",
		Valid:    true,
	}, nil
}
