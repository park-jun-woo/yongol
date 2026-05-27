import { Module } from '@nestjs/common';
import { PrismaModule } from './prisma/prisma.module';
import { AuthzModule } from './authz/authz.module';
import { BillingModule } from './billing/billing.module';
import { GeocodingModule } from './geocoding/geocoding.module';
import { QueueModule } from './queue/queue.module';
import { ReportModule } from './report/report.module';
import { SessionModule } from './session/session.module';
import { VersionModule } from './version/version.module';
import { WebhookdeliveryModule } from './webhookdelivery/webhookdelivery.module';
import { WorkerModule } from './worker/worker.module';
import { AuditModule } from './audit/audit.module';
import { AuthModule } from './auth/auth.module';
import { DashboardModule } from './dashboard/dashboard.module';
import { ExecutionModule } from './execution/execution.module';
import { OrganizationModule } from './organization/organization.module';
import { ScheduleModule } from './schedule/schedule.module';
import { TemplateModule } from './template/template.module';
import { WebhookModule } from './webhook/webhook.module';
import { WorkflowModule } from './workflow/workflow.module';

@Module({
  imports: [
    PrismaModule,
    AuthzModule,
    BillingModule,
    GeocodingModule,
    QueueModule,
    ReportModule,
    SessionModule,
    VersionModule,
    WebhookdeliveryModule,
    WorkerModule,
    AuditModule,
    AuthModule,
    DashboardModule,
    ExecutionModule,
    OrganizationModule,
    ScheduleModule,
    TemplateModule,
    WebhookModule,
    WorkflowModule,
  ],
})
export class AppModule {}
