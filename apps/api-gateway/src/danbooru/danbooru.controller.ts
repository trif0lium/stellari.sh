import { Controller, Get, ParseArrayPipe, Query } from '@nestjs/common';
import { Observable } from 'rxjs';
import { PostsResponse } from 'src/generated/apis/danbooru';
import { DanbooruService } from './danbooru.service';

@Controller('danbooru')
export class DanbooruController {
  constructor(private readonly danbooruService: DanbooruService) {}

  @Get('/')
  getPosts(
    @Query('tags', new ParseArrayPipe({ items: String, separator: ',' })) tags: string[],
  ): Observable<PostsResponse> {
    return this.danbooruService.getPosts(tags);
  }
}
