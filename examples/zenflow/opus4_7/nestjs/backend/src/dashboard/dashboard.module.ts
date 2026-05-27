import { Module } from '@nestjs/common';
import { PrismaModule } from '../../prisma/prisma.module';
import { GetDashboardController } from './getDashboard.controller';
import { GetDashboardService } from './getDashboard.service';

@Module({
  imports: [
    PrismaModule,
  ],
  controllers: [
    GetDashboardController,
  ],
  providers: [
    GetDashboardService,
  ],
})
export class DashboardModule {}
