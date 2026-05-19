package workflow

// @func matchMember
// @description Match workflow to best member based on trigger event. Returns zero-UUID if no match.

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
		// Simulate match found
		return MatchMemberResponse{
			MemberID:   "40000000-0000-0000-0000-000000000004",
			Confidence: "high",
		}, nil
	}
	// No match — return zero UUID sentinel
	return MatchMemberResponse{
		MemberID:   "00000000-0000-0000-0000-000000000000",
		Confidence: "none",
	}, nil
}
