//ff:func feature=external type=generator control=sequence
//ff:what OpenAPI 소스를 읽어 Go 모델 파일을 생성하고 outputDir에 저장한다
package external

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Generate reads an OpenAPI source (file path or URL), generates a Go model file,
// and writes it to outputDir.
func Generate(source, outputDir string) error {
	data, err := readSource(source)
	if err != nil {
		return fmt.Errorf("read source: %w", err)
	}

	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(data)
	if err != nil {
		return fmt.Errorf("parse OpenAPI: %w", err)
	}
	if err := doc.Validate(loader.Context); err != nil {
		return fmt.Errorf("validate OpenAPI: %w", err)
	}

	serviceName := inferServiceName(source, doc)
	code, err := generateCode(serviceName, doc)
	if err != nil {
		return fmt.Errorf("generate code: %w", err)
	}

	outFile := filepath.Join(outputDir, strings.ToLower(serviceName)+".go")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}
	if err := os.WriteFile(outFile, code, 0644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	fmt.Printf("✓ gen-model   %s → %s\n", serviceName, outFile)
	return nil
}
