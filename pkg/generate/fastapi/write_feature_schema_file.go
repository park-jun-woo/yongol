//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeFeatureSchemaFile — 단일 feature 의 Pydantic 스키마를 파일로 출력 (빈 내용 시 skip)

package fastapi

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeFeatureSchemaFile renders the Pydantic schemas for one feature and
// writes them to <schemasDir>/<feature>.py. No file is written when the
// feature has no request-body schemas.
func writeFeatureSchemaFile(schemasDir, feature string, plans []*ir.ServicePlan) error {
	content := renderFeatureSchemas(plans)
	if content == "" {
		return nil
	}
	path := filepath.Join(schemasDir, feature+".py")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write schema %s: %w", feature, err)
	}
	return nil
}
