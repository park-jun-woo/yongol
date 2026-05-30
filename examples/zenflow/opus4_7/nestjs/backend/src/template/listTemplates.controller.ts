import {
  Controller,
  Get,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { ListTemplatesService } from './listTemplates.service';

@Controller('templates')
export class ListTemplatesController {
  constructor(private readonly service: ListTemplatesService) {}

  @Get('')
  async listTemplates(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.listTemplates(params, body, req.user);
  }
}
