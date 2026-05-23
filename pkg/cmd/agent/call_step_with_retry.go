//ff:func feature=agent type=helper control=sequence
//ff:what callStepWithRetry — LLM 호출 후 실패 시 1회 재시도

package agent

import "fmt"

func callStepWithRetry(cfg Config, userPrompt string) (string, error) {
	numCtx := len(splitSystemPrompt) + len(userPrompt) + 2048

	reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, splitSystemPrompt, userPrompt, numCtx)
	if err != nil {
		reply, err = llmCallWithNumCtx(cfg.Backend, cfg.Model, splitSystemPrompt, userPrompt, numCtx)
		if err != nil {
			return "", err
		}
	}

	content := stripCodeBlock(reply)
	if content == "" {
		reply, err = llmCallWithNumCtx(cfg.Backend, cfg.Model, splitSystemPrompt, userPrompt, numCtx)
		if err != nil {
			return "", err
		}
		content = stripCodeBlock(reply)
		if content == "" {
			return "", fmt.Errorf("empty LLM response after retry")
		}
	}

	return content, nil
}
