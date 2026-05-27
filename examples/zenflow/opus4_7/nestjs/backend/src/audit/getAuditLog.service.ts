import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class GetAuditLogService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async getAuditLog(params: any, user?: any): Promise<any> {
    const owner = await tx.audit_logs.findUnique({
      where: { id: params.id },
      select: { org_id: true },
    });
    await this.authz.check({
      action: 'GetAuditLog',
      resource: 'audit_log',
      ResourceID: params.id,
      resourceId: String(params.id),
      owners: { audit_logs: { org_id: owner?.org_id } },
    });
    const audit_log = await this.prisma.auditLog.findUnique({ where: { id: params.id } });
    if (!audit_log) {
      throw new HttpException('Audit log not found', HttpStatus.NOT_FOUND);
    }
    return {
      audit_log: audit_log,
    };
  }
}
