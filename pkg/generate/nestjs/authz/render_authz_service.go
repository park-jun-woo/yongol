//ff:func feature=gen-nestjs type=generator control=sequence
//ff:what RenderAuthzService — AuthzService stub TypeScript 소스 생성

package authz

// RenderAuthzService produces the NestJS authz service stub. The check method
// provides the call structure; actual OPA evaluation is user-implemented.
func RenderAuthzService() string {
	return `import { Injectable, ForbiddenException, Logger } from '@nestjs/common';

export interface AuthzInput {
  action: string;
  resource: string;
  [key: string]: any;
}

@Injectable()
export class AuthzService {
  private readonly logger = new Logger(AuthzService.name);

  async check(input: AuthzInput): Promise<void> {
    this.logger.debug(` + "`" + `authz check: ${input.action} on ${input.resource}` + "`" + `);
    // TODO: integrate OPA/Rego policy evaluation
    // throw new ForbiddenException('access denied') when policy rejects
  }
}
`
}
