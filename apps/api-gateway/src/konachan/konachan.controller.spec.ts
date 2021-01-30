import { Test, TestingModule } from '@nestjs/testing';
import { KonachanController } from './konachan.controller';

describe('KonachanController', () => {
  let controller: KonachanController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [KonachanController],
    }).compile();

    controller = module.get<KonachanController>(KonachanController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
