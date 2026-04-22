//ff:func feature=chain type=formatter control=selection
//ff:what ownership 타입별 표시용 아이콘 문자열 반환
package chain

// ownershipIcon returns the display icon for an ownership type.
func ownershipIcon(ownership string) string {
	switch ownership {
	case "preserve":
		return "preserve"
	case "gen":
		return "gen"
	default:
		return ""
	}
}
