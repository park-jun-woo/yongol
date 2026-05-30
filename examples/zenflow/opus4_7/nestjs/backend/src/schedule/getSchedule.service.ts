import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';
import { ScheduleService } from '../../schedule/schedule.service';
import { SessionService } from '../../session/session.service';

@Injectable()
export class GetScheduleService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
    private readonly scheduleService: ScheduleService,
    private readonly sessionService: SessionService,
  ) {}

  async getSchedule(params: any, body: any, user?: any): Promise<any> {
    await this.authz.check({
      action: 'GetSchedule',
      resource: 'workflow',
      ResourceID: params.id,
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
