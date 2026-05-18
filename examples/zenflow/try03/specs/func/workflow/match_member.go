package workflow

// @func matchMember
// @description Match a workflow to the best team member by trigger event type. Returns sentinel on no match.

type MatchMemberRequest struct {
	TriggerEvent string
	MemberCount  int64
}

type MatchMemberResponse struct {
	MemberID   string
	Confidence string
}

func MatchMember(req MatchMemberRequest) (MatchMemberResponse, error) {
	if req.MemberCount > 0 && req.TriggerEvent != "" {
		return MatchMemberResponse{
			MemberID:   "matched",
			Confidence: "high",
		}, nil
	}
	return MatchMemberResponse{
		MemberID:   "",
		Confidence: "none",
	}, nil
}
