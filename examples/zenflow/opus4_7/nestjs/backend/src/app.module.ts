import { Module } from '@nestjs/common';
import { PrismaModule } from './prisma/prisma.module';
import { DashboardModule } from './dashboard/dashboard.module';
import { ExecutionModule } from './execution/execution.module';
import { OrganizationModule } from './organization/organization.module';
import { ScheduleModule } from './schedule/schedule.module';
import { WebhookModule } from './webhook/webhook.module';
import { AuditModule } from './audit/audit.module';
import { AuthModule } from './auth/auth.module';
import { TemplateModule } from './template/template.module';
import { WorkflowModule } from './workflow/workflow.module';

@Module({
  imports: [
    PrismaModule,
    DashboardModule,
    ExecutionModule,
    OrganizationModule,
    ScheduleModule,
    WebhookModule,
    AuditModule,
    AuthModule,
    TemplateModule,
    WorkflowModule,
  ],
})
export class AppModule {}
