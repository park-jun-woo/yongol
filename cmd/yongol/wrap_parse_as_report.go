//ff:func feature=cli type=helper control=sequence
//ff:what wrapParseAsReport — parse diagnostics를 validate.Report 한 step("parse")으로 감싸 JSON 렌더링 가능하게 반환
package main

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func wrapParseAsReport(diags []diagnostic.Diagnostic) *validate.Report {
	return &validate.Report{
		Steps: []validate.StepResult{
			{
				Name:        "parse",
				Status:      validate.StatusFail,
				Summary:     fmt.Sprintf("%d parse error(s)", len(diags)),
				Diagnostics: diags,
			},
		},
	}
}
