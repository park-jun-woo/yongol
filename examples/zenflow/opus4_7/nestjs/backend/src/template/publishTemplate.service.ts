import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';

@Injectable()
export class PublishTemplateService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async publishTemplate(params: any, body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      await this.authz.check({
        action: 'PublishTemplate',
        resource: 'template',
      });
      const existing = await tx.template.findUnique({ where: { source_workflow_id: params.source_workflow_id } });
      if (existing) {
        throw new HttpException('Already published', HttpStatus.CONFLICT);
      }
      const template = await tx.template.create({ data: { category: params.category, description: params.description, org_id: user.org_id, source_workflow_id: params.source_workflow_id, title: params.title } });
      return {
        template: template,
      };
    });
  }
}
