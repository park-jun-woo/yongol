import {
  Controller,
  Post,
  Param,
  Body,
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
    @Body() body: any,
  ) {
    return this.service.executeWorkflow(params, body, req.user);
  }
}
