import { Module } from '@nestjs/common';
import { PrismaModule } from '../prisma/prisma.module';
import { AuthzModule } from '../authz/authz.module';
import { DeleteScheduleController } from './deleteSchedule.controller';
import { DeleteScheduleService } from './deleteSchedule.service';
import { GetScheduleController } from './getSchedule.controller';
import { GetScheduleService } from './getSchedule.service';
import { SetScheduleController } from './setSchedule.controller';
import { SetScheduleService } from './setSchedule.service';

@Module({
  imports: [
    PrismaModule,
    AuthzModule,
  ],
  controllers: [
    DeleteScheduleController,
    GetScheduleController,
    SetScheduleController,
  ],
  providers: [
    DeleteScheduleService,
    GetScheduleService,
    SetScheduleService,
  ],
})
export class ScheduleModule {}
