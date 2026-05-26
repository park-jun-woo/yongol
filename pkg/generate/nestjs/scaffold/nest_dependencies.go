//ff:func feature=gen-nestjs type=util control=sequence
//ff:what nestDependencies — NestJS 프로젝트 런타임 의존성 맵

package scaffold

// nestDependencies returns the runtime dependencies for a NestJS project.
func nestDependencies() map[string]string {
	return map[string]string{
		"@nestjs/common":              "^10.0.0",
		"@nestjs/core":                "^10.0.0",
		"@nestjs/platform-express":    "^10.0.0",
		"@prisma/client":              "^5.0.0",
		"class-transformer":           "^0.5.1",
		"class-validator":             "^0.14.0",
		"reflect-metadata":            "^0.1.13",
		"rxjs":                        "^7.8.0",
	}
}
