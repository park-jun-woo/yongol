//ff:func feature=gen-nestjs type=generator control=sequence
//ff:what RenderAuthzModule — AuthzModule TypeScript 소스 생성

package authz

// RenderAuthzModule produces the NestJS authz module file content.
func RenderAuthzModule() string {
	return `import { Global, Module } from '@nestjs/common';
import { AuthzService } from './authz.service';

@Global()
@Module({
  providers: [AuthzService],
  exports: [AuthzService],
})
export class AuthzModule {}
`
}
