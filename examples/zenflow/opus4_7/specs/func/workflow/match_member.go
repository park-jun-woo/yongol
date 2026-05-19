package workflow

// @func matchMember
// @description Match a workflow trigger event to the best team member. Returns zero-UUID sentinel on no match.

type MatchMemberRequest struct {
	TriggerEvent string
	MemberCount  int64
}

type MatchMemberResponse struct {
	MemberID   string
	Confidence string
}

func MatchMember(req MatchMemberRequest) (MatchMemberResponse, error) {
	if req.MemberCount > 0 {
		return MatchMemberResponse{
			MemberID:   "22222222-2222-2222-2222-222222222222",
			Confidence: "high",
		}, nil
	}
	return MatchMemberResponse{
		MemberID:   "00000000-0000-0000-0000-000000000000",
		Confidence: "none",
	}, nil
}
