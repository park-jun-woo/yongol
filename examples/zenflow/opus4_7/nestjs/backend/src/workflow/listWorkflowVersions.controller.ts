import {
  Controller,
  Get,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { ListWorkflowVersionsService } from './listWorkflowVersions.service';

@Controller('workflows')
export class ListWorkflowVersionsController {
  constructor(private readonly service: ListWorkflowVersionsService) {}

  @Get(':id/versions')
  async listWorkflowVersions(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.listWorkflowVersions(params, body, req.user);
  }
}
