import { Module } from '@nestjs/common';
import { PrismaModule } from '../prisma/prisma.module';
import { AuthzModule } from '../authz/authz.module';
import { GetDashboardController } from './getDashboard.controller';
import { GetDashboardService } from './getDashboard.service';

@Module({
  imports: [
    PrismaModule,
    AuthzModule,
  ],
  controllers: [
    GetDashboardController,
  ],
  providers: [
    GetDashboardService,
  ],
})
export class DashboardModule {}
