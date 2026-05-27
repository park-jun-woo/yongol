import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';
import { DashboardService } from '../dashboard/dashboard.service';

@Injectable()
export class GetExecutionDetailService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
    private readonly dashboardService: DashboardService,
  ) {}

  async getExecutionDetail(params: any, user?: any): Promise<any> {
    const owner = await this.prisma.execution_log.findUnique({
      where: { id: params.id },
      select: { org_id: true },
    });
    await this.authz.check({
      action: 'GetExecutionDetail',
      resource: 'execution_log',
      resourceId: String(params.id),
      owners: { execution_log: { org_id: owner?.org_id } },
    });
    const log = await this.prisma.executionLog.findUnique({ where: { id: params.id } });
    if (!log) {
      throw new HttpException('Execution log not found', HttpStatus.NOT_FOUND);
    }
    const wf = await this.prisma.workflow.findUnique({ where: { id: log.workflow_id } });
    if (!wf) {
      throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
    }
    const org = await this.prisma.organization.findUnique({ where: { id: log.org_id } });
    if (!org) {
      throw new HttpException('Organization not found', HttpStatus.NOT_FOUND);
    }
    const detail = await this.dashboardService.buildExecutionDetail(log.credits_spent, 'now', 'now', 'now', org.name, log.status, 'now', wf.title);
    return {
      detail: detail,
    };
  }
}
