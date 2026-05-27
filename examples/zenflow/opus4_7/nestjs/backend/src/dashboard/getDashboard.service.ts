import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';
import { DashboardService } from '../../dashboard/dashboard.service';

@Injectable()
export class GetDashboardService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
    private readonly dashboardService: DashboardService,
  ) {}

  async getDashboard(params: any, body: any, user?: any): Promise<any> {
    await this.authz.check({
      action: 'GetDashboard',
      resource: 'organization',
    });
    const org = await this.prisma.organization.findUnique({ where: { id: user.org_id } });
    if (!org) {
      throw new HttpException('Organization not found', HttpStatus.NOT_FOUND);
    }
    const summary = await this.dashboardService.summarize(org.credits_balance, org.name, org.plan_type);
    return {
      summary: summary,
    };
  }
}
