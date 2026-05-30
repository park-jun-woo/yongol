import {
  Controller,
  Post,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { CreateWorkflowService } from './createWorkflow.service';

@Controller('workflows')
export class CreateWorkflowController {
  constructor(private readonly service: CreateWorkflowService) {}

  @Post('')
  async createWorkflow(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.createWorkflow(params, body, req.user);
  }
}
