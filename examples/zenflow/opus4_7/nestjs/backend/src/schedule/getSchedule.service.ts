import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';
import { ScheduleService } from '../schedule/schedule.service';
import { SessionService } from '../session/session.service';

@Injectable()
export class GetScheduleService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
    private readonly scheduleService: ScheduleService,
    private readonly sessionService: SessionService,
  ) {}

  async getSchedule(params: any, user?: any): Promise<any> {
    const owner = await this.prisma.workflow.findUnique({
      where: { id: params.id },
      select: { org_id: true },
    });
    await this.authz.check({
      action: 'GetSchedule',
      resource: 'workflow',
      resourceId: String(params.id),
      owners: { workflow: { org_id: owner?.org_id } },
    });
    const wf = await this.prisma.workflow.findUnique({ where: { id: params.id } });
    if (!wf) {
      throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
    }
    const keyResult = await this.scheduleService.buildKey(params.id);
    const sessionResult = await this.sessionService.get(keyResult.key);
    return {
      cron: sessionResult.Value,
    };
  }
}
