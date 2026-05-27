import {
  Controller,
  Get,
  Param,
  Req,
} from '@nestjs/common';
import { GetWorkflowService } from './getWorkflow.service';

@Controller('workflows')
export class GetWorkflowController {
  constructor(private readonly service: GetWorkflowService) {}

  @Get(':id')
  async getWorkflow(
    @Req() req: any,
    @Param() params: any,
  ) {
    return this.service.getWorkflow(params, req.user);
  }
}
