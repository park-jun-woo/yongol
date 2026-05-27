import {
  Controller,
  Post,
  Param,
  Req,
} from '@nestjs/common';
import { ExecuteWorkflowService } from './executeWorkflow.service';

@Controller('workflows')
export class ExecuteWorkflowController {
  constructor(private readonly service: ExecuteWorkflowService) {}

  @Post(':id/execute')
  async executeWorkflow(
    @Req() req: any,
    @Param() params: any,
  ) {
    return this.service.executeWorkflow(params, req.user);
  }
}
