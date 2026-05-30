import {
  Controller,
  Get,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { ListWorkflowsService } from './listWorkflows.service';

@Controller('workflows')
export class ListWorkflowsController {
  constructor(private readonly service: ListWorkflowsService) {}

  @Get('')
  async listWorkflows(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.listWorkflows(params, body, req.user);
  }
}
