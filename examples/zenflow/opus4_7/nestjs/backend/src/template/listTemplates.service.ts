import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';

@Injectable()
export class ListTemplatesService {
  constructor(
    private readonly prisma: PrismaService,
  ) {}

  async listTemplates(params: any, body: any, user?: any): Promise<any> {
    const items = await this.prisma.template.findMany({ where: { category: params.category, cursor: params.cursor, per_page: params.per_page } });
    return {
      items: items,
    };
  }
}
