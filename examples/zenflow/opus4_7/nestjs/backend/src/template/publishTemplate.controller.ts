import {
  Controller,
  Post,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { PublishTemplateService } from './publishTemplate.service';

@Controller('templates')
export class PublishTemplateController {
  constructor(private readonly service: PublishTemplateService) {}

  @Post('')
  async publishTemplate(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.publishTemplate(params, body, req.user);
  }
}
