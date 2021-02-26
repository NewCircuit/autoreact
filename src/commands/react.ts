import { TextChannel } from 'discord.js';
import { Command, CommandoClient, CommandoMessage } from 'discord.js-commando';
import { CONFIG } from '../global';

export class React extends Command {
  constructor(client: CommandoClient) {
    super(client, {
      name: 'react',
      description: 'reacts to all messages in a channel',
      group: 'staff',
      memberName: 'react',
      userPermissions: ['MANAGE_CHANNELS'],
    });
  }

  public async run(msg: CommandoMessage): Promise<null> {
    const channel = msg.channel as TextChannel;
    const messages = await channel.messages.fetch({ limit: 100 });
    const msgs = messages.values();
    const tasks = [];

    let result = msgs.next();
    while (!result.done) {
      const message = result.value;
      for (let i = 0; i < CONFIG.emojis.length; i += 1) {
        tasks.push(message.react(CONFIG.emojis[i]));
      }
      result = msgs.next();
    }

    await Promise.all(tasks);
    const botmsg = await msg.reply('all done, autodeletes in 1 minute');
    try {
      await botmsg.delete({ timeout: 60000 });
    } catch (_) {
      return null;
    }
    return null;
  }
}
