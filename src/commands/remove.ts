import { TextChannel } from 'discord.js';
import { Command, CommandoClient, CommandoMessage } from 'discord.js-commando';
import Config from '../config';
import { CONFIG } from '../global';

type Args={
    channel:TextChannel
}

export class Remove extends Command {
  constructor(client: CommandoClient) {
    super(client, {
      name: 'remove',
      description: 'removes a class',
      group: 'staff',
      memberName: 'remove',
      args: [
        {
          key: 'channel',
          prompt: 'please provide a channel',
          type: 'text-channel',
        },
      ],
      userPermissions: ['MANAGE_CHANNELS'],
    });
  }

  public async run(msg: CommandoMessage, args: Args): Promise<null> {
    const { id } = args.channel;
    let botmsg;
    if (CONFIG.whitelist.includes(msg.channel.id)) {
      const i = CONFIG.whitelist.indexOf(id);
      const filtered = CONFIG.whitelist.splice(i, 1);
      CONFIG.whitelist = filtered;
      Config.setConfig(CONFIG);
      botmsg = await msg.reply(`The channel with ID: ${id} is removed.`);
    } else {
      botmsg = await msg.reply(`The channel with ID: ${id} is not whitelisted.`);
    }

    await botmsg.delete({ timeout: 5000 });
    return null;
  }
}
