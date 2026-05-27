import {
  Controller,
  Get,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { ListWebhooksService } from './listWebhooks.service';

@Controller('webhooks')
export class ListWebhooksController {
  constructor(private readonly service: ListWebhooksService) {}

  @Get('')
  async listWebhooks(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.listWebhooks(params, body, req.user);
  }
}
