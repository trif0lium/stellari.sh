import { Test, TestingModule } from '@nestjs/testing';
import { DanbooruService } from './danbooru.service';

describe('DanbooruService', () => {
  let service: DanbooruService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [DanbooruService],
    }).compile();

    service = module.get<DanbooruService>(DanbooruService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
