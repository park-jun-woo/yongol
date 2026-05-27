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
    const owner = await tx.audit_logs.findUnique({
      where: { id: params.id },
      select: { org_id: true },
    });
    await this.authz.check({
      action: 'ListAuditLogs',
      resource: 'audit_log',
      resourceId: String(params.id),
      owners: { audit_logs: { org_id: owner?.org_id } },
    });
    const items = await this.prisma.auditLog.findMany({ where: { filter_action: query.action, org_id: user.org_id, page_offset: query.page, per_page: query.per_page } });
    const total = await this.prisma.auditLog.findUnique({ where: { filter_action: query.action, org_id: user.org_id } });
    return {
      items: items,
      total: total,
    };
  }
}
