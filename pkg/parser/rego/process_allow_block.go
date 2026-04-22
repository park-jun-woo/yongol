//ff:func feature=policy type=parser control=sequence
//ff:what processAllowBlock — allow 블록 하나에서 AllowRule 추출
package rego

func processAllowBlock(block string) (AllowRule, bool) {
	rule := AllowRule{}
	if m := reAction.FindStringSubmatch(block); m != nil {
		if m[1] != "" {
			rule.Actions = []string{m[1]}
		} else if m[2] != "" {
			rule.Actions = parseActionSet(m[2])
		}
	}
	if m := reResource.FindStringSubmatch(block); m != nil {
		rule.Resource = m[1]
	}
	rule.UsesOwner = reOwnerRef.MatchString(block)
	if m := reRoleRef.FindStringSubmatch(block); m != nil {
		rule.UsesRole = true
		rule.RoleValue = m[1]
	}
	if len(rule.Actions) > 0 && rule.Resource != "" {
		return rule, true
	}
	return rule, false
}
