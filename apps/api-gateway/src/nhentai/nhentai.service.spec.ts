import { Test, TestingModule } from '@nestjs/testing';
import { NhentaiService } from './nhentai.service';

describe('NhentaiService', () => {
  let service: NhentaiService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [NhentaiService],
    }).compile();

    service = module.get<NhentaiService>(NhentaiService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
