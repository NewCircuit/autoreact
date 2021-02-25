import { Message, TextChannel } from 'discord.js';
import {Command,CommandoClient,CommandoMessage} from 'discord.js-commando';
import { config } from 'process';
import Config from '../config';
import { CONFIG } from '../global';

type Args= {
    channel:TextChannel
}

export class Add extends Command{
    constructor(client: CommandoClient) {
        super(client,{
            name: 'add',
            description: 'adds a channel',
            group: 'staff',
            memberName: 'add',
            args: [
                {
                    key: 'channel',
                    prompt: 'please provide a channel',
                    type: 'text-channel'
                }
            ]
        });
        
    }
    public async run(msg: CommandoMessage,args: Args): Promise<null>{
        const id = args.channel.id;
        let botmsg;
        await msg.delete();
        if(!CONFIG.whitelist.includes(msg.channel.id)){
            CONFIG.whitelist.push(id);
            Config.setConfig(CONFIG);
            botmsg = await msg.reply('youre id is: '+id);
        }
        else{
            botmsg = await msg.reply('ID: '+ id + ' exists already');
        }
        
        await botmsg.delete({timeout: 5000});
        return null;
    }
}