//ff:func feature=agent type=helper control=sequence
//ff:what buildSystemPrompt — docs 섹션 + 레이어 예시로 system prompt 구성

package agent

// buildSystemPrompt returns the system prompt with docs sections and layer example.
func buildSystemPrompt(l layer, diagMsgs []string) string {
	base := "You fix yongol SSOT files. Output ONLY the corrected file content. No explanations. No markdown.\n\n"

	docSection := searchDocs(l, diagMsgs)
	if docSection != "" {
		base += docSection + "\n\n"
	}

	base += "Example for " + layerName(l) + ":\n" + layerExample(l)
	return base
}
