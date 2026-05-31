//ff:func feature=gen-fastapi type=generator control=iteration dimension=1
//ff:what writeSchemas — ServicePlan.BodyFields 기반 Pydantic BaseModel 스키마 파일 생성

package fastapi

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeSchemas generates Pydantic BaseModel schema files from ServicePlan
// BodyFields for each feature that has at least one POST/PUT/PATCH endpoint
// with a request body.
func writeSchemas(plansByFeature map[string][]*ir.ServicePlan, appDir string) error {
	schemasDir := filepath.Join(appDir, "schemas")
	if err := os.MkdirAll(schemasDir, 0o755); err != nil {
		return fmt.Errorf("mkdir schemas: %w", err)
	}

	for feature, plans := range plansByFeature {
		if err := writeFeatureSchemaFile(schemasDir, feature, plans); err != nil {
			return err
		}
	}
	return nil
}
