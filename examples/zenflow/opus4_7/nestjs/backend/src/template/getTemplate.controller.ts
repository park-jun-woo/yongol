import {
  Controller,
  Get,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { GetTemplateService } from './getTemplate.service';

@Controller('templates')
export class GetTemplateController {
  constructor(private readonly service: GetTemplateService) {}

  @Get(':id')
  async getTemplate(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.getTemplate(params, body, req.user);
  }
}
