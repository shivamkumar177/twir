import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { config } from '@tsuwari/config';

import { DefaultCommandsCreatorService } from './default-commands-creator.service.js';

@Module({
  imports: [
    ClientsModule.register([
      { name: 'NATS', transport: Transport.NATS, options: { servers: [config.NATS_URL] } },
    ]),
  ],
  providers: [DefaultCommandsCreatorService],
})
export class DefaultCommandsCreatorModule { }