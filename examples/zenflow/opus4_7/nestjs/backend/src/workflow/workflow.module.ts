import { Module } from '@nestjs/common';
import { PrismaModule } from '../../prisma/prisma.module';
import { QueueModule } from '../../queue/queue.module';
import { ActivateWorkflowController } from './activateWorkflow.controller';
import { ActivateWorkflowService } from './activateWorkflow.service';
import { AddActionController } from './addAction.controller';
import { AddActionService } from './addAction.service';
import { ArchiveWorkflowController } from './archiveWorkflow.controller';
import { ArchiveWorkflowService } from './archiveWorkflow.service';
import { AutoAssignWorkflowController } from './autoAssignWorkflow.controller';
import { AutoAssignWorkflowService } from './autoAssignWorkflow.service';
import { CreateWorkflowController } from './createWorkflow.controller';
import { CreateWorkflowService } from './createWorkflow.service';
import { CreateWorkflowVersionController } from './createWorkflowVersion.controller';
import { CreateWorkflowVersionService } from './createWorkflowVersion.service';
import { ExecuteWithReportController } from './executeWithReport.controller';
import { ExecuteWithReportService } from './executeWithReport.service';
import { ExecuteWorkflowController } from './executeWorkflow.controller';
import { ExecuteWorkflowService } from './executeWorkflow.service';
import { GetWorkflowController } from './getWorkflow.controller';
import { GetWorkflowService } from './getWorkflow.service';
import { ListExecutionLogsController } from './listExecutionLogs.controller';
import { ListExecutionLogsService } from './listExecutionLogs.service';
import { ListWorkflowVersionsController } from './listWorkflowVersions.controller';
import { ListWorkflowVersionsService } from './listWorkflowVersions.service';
import { ListWorkflowsController } from './listWorkflows.controller';
import { ListWorkflowsService } from './listWorkflows.service';
import { PauseWorkflowController } from './pauseWorkflow.controller';
import { PauseWorkflowService } from './pauseWorkflow.service';
import { SaveWorkflowActionsController } from './saveWorkflowActions.controller';
import { SaveWorkflowActionsService } from './saveWorkflowActions.service';

@Module({
  imports: [
    PrismaModule,
    QueueModule,
  ],
  controllers: [
    ActivateWorkflowController,
    AddActionController,
    ArchiveWorkflowController,
    AutoAssignWorkflowController,
    CreateWorkflowController,
    CreateWorkflowVersionController,
    ExecuteWithReportController,
    ExecuteWorkflowController,
    GetWorkflowController,
    ListExecutionLogsController,
    ListWorkflowVersionsController,
    ListWorkflowsController,
    PauseWorkflowController,
    SaveWorkflowActionsController,
  ],
  providers: [
    ActivateWorkflowService,
    AddActionService,
    ArchiveWorkflowService,
    AutoAssignWorkflowService,
    CreateWorkflowService,
    CreateWorkflowVersionService,
    ExecuteWithReportService,
    ExecuteWorkflowService,
    GetWorkflowService,
    ListExecutionLogsService,
    ListWorkflowVersionsService,
    ListWorkflowsService,
    PauseWorkflowService,
    SaveWorkflowActionsService,
  ],
})
export class WorkflowModule {}
