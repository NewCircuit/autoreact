# Auto React
This bot auto-reacts to specific channels

## Requirements
 * [NodeJS](https://nodejs.org/)

## Setup

### Configure the Bot
1. Copy `config.example.yml` and change the name to `config.yml`
2. Fill out the `config.yml`

### Run the Bot
```sh
# install dependencies
npm i
# start the bot
npm start
```

## Examples
How to add a channel:
- .add #channelname

How to remove a channel:
- .remove #channelname

How to autoreact the last 100 messages in a channel:
- .react 
