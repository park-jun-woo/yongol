import {
  Controller,
  Get,
  Param,
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
  ) {
    return this.service.getTemplate(params, req.user);
  }
}
