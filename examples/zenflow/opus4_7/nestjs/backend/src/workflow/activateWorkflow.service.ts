import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';
import { BillingService } from '../../billing/billing.service';

@Injectable()
export class ActivateWorkflowService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
    private readonly billingService: BillingService,
  ) {}

  async activateWorkflow(params: any, body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      await this.authz.check({
        action: 'ActivateWorkflow',
        resource: 'workflow',
        ResourceID: params.id,
      });
      const wf = await tx.workflow.findUnique({ where: { id: params.id } });
      if (!wf) {
        throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
      }
      const org = await tx.organization.findUnique({ where: { id: wf.org_id } });
      if (!org) {
        throw new HttpException('Organization not found', HttpStatus.NOT_FOUND);
      }
      if (await this.billingService.isZeroBalance(org.credits_balance)) {
        throw new HttpException('Insufficient credits', HttpStatus.PAYMENT_REQUIRED);
      }
      // @state workflows.ActivateWorkflow — transition guard
      const allowed_ActivateWorkflow: Record<string, string[]> = {
        // TODO: populate from Mermaid stateDiagram 'workflows'
      };
      if (!(allowed_ActivateWorkflow[wf.status] || []).includes('ActivateWorkflow')) {
        throw new HttpException('Cannot activate workflow', HttpStatus.CONFLICT);
      }
      await tx.workflow.update({ where: { id: wf.id }, data: { status: 'active' } });
      const updated = await tx.workflow.findUnique({ where: { id: wf.id } });
      return {
        workflow: updated,
      };
    });
  }
}
