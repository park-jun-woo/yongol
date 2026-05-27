import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';
import { ScheduleService } from '../../schedule/schedule.service';
import { SessionService } from '../../session/session.service';

@Injectable()
export class DeleteScheduleService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
    private readonly scheduleService: ScheduleService,
    private readonly sessionService: SessionService,
  ) {}

  async deleteSchedule(params: any, body: any, user?: any): Promise<any> {
    await this.authz.check({
      action: 'DeleteSchedule',
      resource: 'workflow',
      ResourceID: params.id,
    });
    const wf = await this.prisma.workflow.findUnique({ where: { id: params.id } });
    if (!wf) {
      throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
    }
    const keyResult = await this.scheduleService.buildKey(params.id);
    await this.sessionService.delete(keyResult.key);
  }
}
