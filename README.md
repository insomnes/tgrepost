# Telegram reposting bot
Bot will repost messages from channel to chat, with optional topic setting for super-groups.

Configured through environment variables.

See `.env.dist` for a list of required variables.

## Running
Via docker-compose:
```bash
docker compose -f docker/docker-compose.yml up --build
```
Without docker:
```bash
go build -o repostbot .
./repostbot
```
