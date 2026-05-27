import {
  Controller,
  Delete,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { DeleteWebhookService } from './deleteWebhook.service';

@Controller('webhooks')
export class DeleteWebhookController {
  constructor(private readonly service: DeleteWebhookService) {}

  @Delete(':id')
  async deleteWebhook(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.deleteWebhook(params, body, req.user);
  }
}
