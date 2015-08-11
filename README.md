# Raspberry Bot
This is my own raspberry bot.
The basic idea is that I want to send commands over various transports (telegram, mail, etc..) and expect an action from him.

Here's a list of actions I want it to do:

- [ ] Play a song or multiple songs upon receiving the title, artist, album genre etc..
- [x] Send me his public ip address
- [ ] Open a port on my home router so I can reach him from anywhere and then safely close the port when I finished
- [x] Send me some status informations (Ram, Cpu)
- [ ] Download some torrents upon receiving the title
- [ ] Face recognition

### Build
```
GOOS=linux GOARCH=arm CGO_ENABLED=0 go build .
go build .
```
