import {
  Controller,
  Get,
  Req,
} from '@nestjs/common';
import { ListWebhooksService } from './listWebhooks.service';

@Controller('webhooks')
export class ListWebhooksController {
  constructor(private readonly service: ListWebhooksService) {}

  @Get('')
  async listWebhooks(
    @Req() req: any,
  ) {
    return this.service.listWebhooks(req.user);
  }
}
