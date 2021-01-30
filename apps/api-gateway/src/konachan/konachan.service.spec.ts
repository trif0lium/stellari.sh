import { Test, TestingModule } from '@nestjs/testing';
import { KonachanService } from './konachan.service';

describe('KonachanService', () => {
  let service: KonachanService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [KonachanService],
    }).compile();

    service = module.get<KonachanService>(KonachanService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
