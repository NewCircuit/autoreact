import * as fs from 'fs';
import * as yaml from 'js-yaml';

const configFilePath = './config.yml';
export default class Config {
    public readonly token: string;

    public readonly prefix: string;

    public whitelist: string[];

    public emojis: string[];

    constructor() {
      this.token = '';
      this.prefix = '';
      this.whitelist = [''];
      this.emojis = [''];
    }

    public static getConfig(): Config {
      const data = fs.readFileSync(configFilePath);
      return yaml.load(data.toString()) as Config;
    }

    public static setConfig(data: Config) {
      fs.writeFileSync(configFilePath, yaml.dump(data));
    }
}
