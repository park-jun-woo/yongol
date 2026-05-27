import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class ListAuditLogsService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async listAuditLogs(query: any, user?: any): Promise<any> {
    await this.authz.check({
      action: 'ListAuditLogs',
      resource: 'audit_log',
    });
    const items = await this.prisma.auditLog.findMany({ where: { filter_action: query.action, org_id: user.org_id }, skip: query.page, take: query.per_page });
    const total = await this.prisma.auditLog.count({ where: { filter_action: query.action, org_id: user.org_id } });
    return {
      items: items,
      total: total,
    };
  }
}
