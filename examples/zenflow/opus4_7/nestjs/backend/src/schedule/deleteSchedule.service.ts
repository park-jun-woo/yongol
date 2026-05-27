import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';
import { ScheduleService } from '../schedule/schedule.service';
import { SessionService } from '../session/session.service';

@Injectable()
export class DeleteScheduleService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
    private readonly scheduleService: ScheduleService,
    private readonly sessionService: SessionService,
  ) {}

  async deleteSchedule(params: any, user?: any): Promise<any> {
    const owner = await tx.workflows.findUnique({
      where: { id: params.id },
      select: { org_id: true },
    });
    await this.authz.check({
      action: 'DeleteSchedule',
      resource: 'workflow',
      ResourceID: params.id,
      resourceId: String(params.id),
      owners: { workflows: { org_id: owner?.org_id } },
    });
    const wf = await this.prisma.workflow.findUnique({ where: { id: params.id } });
    if (!wf) {
      throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
    }
    const keyResult = await this.scheduleService.buildKey(params.id);
    await this.sessionService.delete(keyResult.key);
  }
}
