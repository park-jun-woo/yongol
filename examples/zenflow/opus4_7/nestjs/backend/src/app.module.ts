import { Module } from '@nestjs/common';
import { PrismaModule } from './prisma/prisma.module';
import { AuthModule } from './auth/auth.module';
import { ExecutionModule } from './execution/execution.module';
import { ScheduleModule } from './schedule/schedule.module';
import { TemplateModule } from './template/template.module';
import { WebhookModule } from './webhook/webhook.module';
import { AuditModule } from './audit/audit.module';
import { DashboardModule } from './dashboard/dashboard.module';
import { OrganizationModule } from './organization/organization.module';
import { WorkflowModule } from './workflow/workflow.module';

@Module({
  imports: [
    PrismaModule,
    AuthModule,
    ExecutionModule,
    ScheduleModule,
    TemplateModule,
    WebhookModule,
    AuditModule,
    DashboardModule,
    OrganizationModule,
    WorkflowModule,
  ],
})
export class AppModule {}
