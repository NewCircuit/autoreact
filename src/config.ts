import * as fs from 'fs';
import * as yaml from 'js-yaml';
import { stringify } from 'querystring';

export default class Config{
    public readonly token: string;
    public readonly prefix: string;
    public readonly whitelist: string[];

    constructor(){
        this.token = '';
        this.prefix = '';
        this.whitelist = [''];
    }

    public static getConfig(): Config{
        const data = fs.readFileSync('./config.yaml');
        return yaml.load(data.toString()) as Config;
    }
    public static setConfig(data: Config){
        fs.writeFileSync('./config.yaml', yaml.dump(data));
    }
}