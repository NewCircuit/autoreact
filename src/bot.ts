import {Message} from 'discord.js';
import {CommandoClient} from 'discord.js-commando';
import {CONFIG} from './global';

class Bot extends CommandoClient{
    
    constructor() {
        super({
            commandPrefix: CONFIG.prefix
        });
        
    }
    public async start(){
        this.registry
        .registerDefaultTypes({
            textChannel: true
        })
        .registerGroup('staff')
        .registerCommandsIn(`${__dirname}/commands`);
        this.on('message',this.onMessage.bind(this));
        await this.login(CONFIG.token);
    }
    public async onMessage(msg: Message){
        if(CONFIG.whitelist.includes(msg.channel.id)){
            await msg.react('✔');
            await msg.react('😊');
        }
    }
}

const bot = new Bot();
bot.start();