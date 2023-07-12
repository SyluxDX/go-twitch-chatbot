# Go Twicth Chatbot

|             | WebSocket clients               | IRC clients                   |
| ----------- | ------------------------------- | ----------------------------- |
| **SSL**     | wss://irc-ws.chat.twitch.tv:443 | irc://irc.chat.twitch.tv:6697 |
| **Non SSL** | ws://irc-ws.chat.twitch.tv:80   | irc://irc.chat.twitch.tv:6667 |


[Twitch chat docs](https://dev.twitch.tv/docs/irc/)  
[How to build websockets in Go](https://yalantis.com/blog/how-to-build-websockets-in-go/)  
[reader with timeout](https://gist.github.com/hongster/04660a20f2498fb7b680)

[golang console user interface](https://github.com/jroimartin/gocui)


# ping example
```
2023/07/12 08:59:38 PING :tmi.twitch.tv
2023/07/12 09:05:38 ERROR EOF
```
