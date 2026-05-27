import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class CloneTemplateService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async cloneTemplate(params: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      const owner = await tx.templates.findUnique({
        where: { id: params.id },
        select: { org_id: true },
      });
      await this.authz.check({
        action: 'CloneTemplate',
        resource: 'template',
        ResourceID: params.id,
        resourceId: String(params.id),
        owners: { templates: { org_id: owner?.org_id } },
      });
      const tmpl = await tx.template.findUnique({ where: { id: params.id } });
      if (!tmpl) {
        throw new HttpException('Template not found', HttpStatus.NOT_FOUND);
      }
      const sourceWf = await tx.workflow.findUnique({ where: { id: tmpl.id } });
      if (!sourceWf) {
        throw new HttpException('Source workflow not found', HttpStatus.NOT_FOUND);
      }
      const newWf = await tx.workflow.create({ data: { org_id: user.org_id, title: sourceWf.title, trigger_event: sourceWf.trigger_event } });
      await tx.action.update({ where: { id: params.id }, data: { new_workflow_id: newWf.id, source_workflow_id: sourceWf.id } });
      await tx.template.update({ where: { id: tmpl.id }, data: { ...body } });
      return {
        workflow: newWf,
      };
    });
  }
}
