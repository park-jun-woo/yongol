import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class PublishTemplateService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async publishTemplate(body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      const owner = await tx.templates.findUnique({
        where: { id: params.id },
        select: { org_id: true },
      });
      await this.authz.check({
        action: 'PublishTemplate',
        resource: 'template',
        resourceId: String(params.id),
        owners: { templates: { org_id: owner?.org_id } },
      });
      const existing = await tx.template.findUnique({ where: { source_workflow_id: body.source_workflow_id } });
      if (existing) {
        throw new HttpException('Already published', HttpStatus.CONFLICT);
      }
      const template = await tx.template.create({ data: { category: body.category, description: body.description, org_id: user.org_id, source_workflow_id: body.source_workflow_id, title: body.title } });
      return {
        template: template,
      };
    });
  }
}
