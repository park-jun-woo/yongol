import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';
import { BillingService } from '../billing/billing.service';
import { ReportService } from '../report/report.service';
import { WorkerService } from '../worker/worker.service';

@Injectable()
export class ExecuteWithReportService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
    private readonly billingService: BillingService,
    private readonly reportService: ReportService,
    private readonly workerService: WorkerService,
  ) {}

  async executeWithReport(params: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      const owner = await tx.workflows.findUnique({
        where: { id: params.id },
        select: { org_id: true },
      });
      await this.authz.check({
        action: 'ExecuteWorkflow',
        resource: 'workflow',
        ResourceID: params.id,
        resourceId: String(params.id),
        owners: { workflows: { org_id: owner?.org_id } },
      });
      const wf = await tx.workflow.findUnique({ where: { id: params.id } });
      if (!wf) {
        throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
      }
      // @state workflows.ExecuteWorkflow — transition guard
      const allowed_ExecuteWorkflow: Record<string, boolean> = {
        'active': true,
      };
      if (!allowed_ExecuteWorkflow[wf.status]) {
        throw new HttpException('Workflow is not active', HttpStatus.CONFLICT);
      }
      const org = await tx.organization.findUnique({ where: { id: wf.id } });
      if (!org) {
        throw new HttpException('Organization not found', HttpStatus.NOT_FOUND);
      }
      if (await this.billingService.isZeroBalance(org.credits_balance)) {
        throw new HttpException('Insufficient credits', HttpStatus.PAYMENT_REQUIRED);
      }
      const actionResult = await this.workerService.processAction(wf.trigger_event, '');
      await tx.organization.update({ where: { id: org.id }, data: { amount: 1 } });
      const reportResult = await this.reportService.generateReport(1, 'completed', wf.title);
      const log = await tx.executionLog.create({ data: { credits_spent: 1.credits_spent, org_id: wf.org_id, report_key: reportResult.report_key, status: 'completed', workflow_id: wf.workflow_id } });
      return {
        action_result: actionResult,
        execution_log: log,
        report: reportResult,
      };
    });
  }
}
