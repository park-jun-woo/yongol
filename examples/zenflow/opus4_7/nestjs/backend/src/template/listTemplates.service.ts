import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';

@Injectable()
export class ListTemplatesService {
  constructor(
    private readonly prisma: PrismaService,
  ) {}

  async listTemplates(query: any, user?: any): Promise<any> {
    const items = await this.prisma.template.findMany({ where: { category: query.category, cursor: query.cursor, per_page: query.per_page } });
    return {
      items: items,
    };
  }
}
