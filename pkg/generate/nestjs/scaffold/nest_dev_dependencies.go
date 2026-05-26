//ff:func feature=gen-nestjs type=util control=sequence
//ff:what nestDevDependencies — NestJS 프로젝트 개발 의존성 맵

package scaffold

// nestDevDependencies returns the dev dependencies for a NestJS project.
func nestDevDependencies() map[string]string {
	return map[string]string{
		"@nestjs/cli":     "^10.0.0",
		"@nestjs/testing": "^10.0.0",
		"@types/node":     "^20.0.0",
		"@types/jest":     "^29.5.0",
		"jest":            "^29.5.0",
		"prisma":          "^5.0.0",
		"ts-jest":         "^29.1.0",
		"typescript":      "^5.1.0",
		"ts-node":         "^10.9.0",
	}
}
