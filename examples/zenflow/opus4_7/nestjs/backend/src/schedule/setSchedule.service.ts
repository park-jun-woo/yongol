import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';
import { ScheduleService } from '../../schedule/schedule.service';
import { SessionService } from '../../session/session.service';

@Injectable()
export class SetScheduleService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
    private readonly scheduleService: ScheduleService,
    private readonly sessionService: SessionService,
  ) {}

  async setSchedule(params: any, body: any, user?: any): Promise<any> {
    await this.authz.check({
      action: 'SetSchedule',
      resource: 'workflow',
      ResourceID: params.id,
    });
    const wf = await this.prisma.workflow.findUnique({ where: { id: params.id } });
    if (!wf) {
      throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
    }
    const cronResult = await this.scheduleService.parseCron(params.cron);
    const keyResult = await this.scheduleService.buildKey(params.id);
    await this.sessionService.set(keyResult.key, 86400, params.cron);
    return {
      schedule: cronResult,
    };
  }
}
