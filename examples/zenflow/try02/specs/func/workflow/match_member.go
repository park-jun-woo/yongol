package workflow

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/zenflow/internal/db"
)

// @func matchMember
// @error 500
// @description Pure matching logic: given a workflow trigger_event and a list of org members,
// returns the best-matching member ID string, or the zero-UUID sentinel when no match is found.

const zeroUUID = "00000000-0000-0000-0000-000000000000"

// UserListByOrgRow is a type alias for db.UserListByOrgRow so that
// the @call SSaC type inference (bare name) matches this package's type.
type UserListByOrgRow = db.UserListByOrgRow

type MatchMemberRequest struct {
	TriggerEvent string
	Members      []UserListByOrgRow
}

type MatchMemberResponse struct {
	MemberID   string
	Confidence string
}

func uuidToString(b [16]byte) string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

func MatchMember(req MatchMemberRequest) (MatchMemberResponse, error) {
	if len(req.Members) == 0 {
		return MatchMemberResponse{MemberID: zeroUUID, Confidence: "none"}, nil
	}
	event := req.TriggerEvent
	for _, m := range req.Members {
		if strings.HasPrefix(event, "admin.") && m.Role == "admin" {
			return MatchMemberResponse{MemberID: uuidToString(m.ID.Bytes), Confidence: "high"}, nil
		}
		if strings.HasPrefix(event, "user.") && m.Role == "member" {
			return MatchMemberResponse{MemberID: uuidToString(m.ID.Bytes), Confidence: "high"}, nil
		}
	}
	return MatchMemberResponse{MemberID: zeroUUID, Confidence: "none"}, nil
}
