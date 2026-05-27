import {
  Controller,
  Get,
  Param,
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
  ) {
    return this.service.listWorkflowVersions(params, req.user);
  }
}
