//ff:func feature=gen-react type=generator control=sequence
//ff:what writeOpenapiTsStub — openapi-typescript 부재 시 fallback types/api.d.ts 기록

package react

import "os"

// writeOpenapiTsStub writes a minimal types/api.d.ts that makes TypeScript
// compile (barely) while surfacing the failure reason in a comment. This
// is strictly a fallback so AI iteration loops aren't blocked by a missing
// node_modules install; the generate command still returns a non-nil error
// so automation surfaces the issue.
func writeOpenapiTsStub(destPath string, reason error) {
	content := "// openapi-typescript could not run: " + reason.Error() + "\n" +
		"// Install it in your frontend project: npm install --save-dev openapi-typescript\n" +
		"export type paths = Record<string, any>\n"
	_ = os.WriteFile(destPath, []byte(content), 0o644)
}
