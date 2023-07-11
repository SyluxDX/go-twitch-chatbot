# Go Twicth Chatbot

|             | WebSocket clients               | IRC clients                   |
| ----------- | ------------------------------- | ----------------------------- |
| **SSL**     | wss://irc-ws.chat.twitch.tv:443 | irc://irc.chat.twitch.tv:6697 |
| **Non SSL** | ws://irc-ws.chat.twitch.tv:80   | irc://irc.chat.twitch.tv:6667 |


[Twitch chat docs](https://dev.twitch.tv/docs/irc/)  
[How to build websockets in Go](https://yalantis.com/blog/how-to-build-websockets-in-go/)  
[reader with timeout](https://gist.github.com/hongster/04660a20f2498fb7b680)



# Start of connection
```
PASS justinfan6493
NICK justinfan6493
JOIN #syluxdx

2023/07/11 09:14:12 :tmi.twitch.tv 001 justinfan6493 :Welcome, GLHF!
2023/07/11 09:14:12 :tmi.twitch.tv 002 justinfan6493 :Your host is tmi.twitch.tv
2023/07/11 09:14:12 :tmi.twitch.tv 003 justinfan6493 :This server is rather new
2023/07/11 09:14:12 :tmi.twitch.tv 004 justinfan6493 :-
2023/07/11 09:14:12 :tmi.twitch.tv 375 justinfan6493 :-
2023/07/11 09:14:12 :tmi.twitch.tv 372 justinfan6493 :You are in a maze of twisty passages, all alike.
2023/07/11 09:14:12 :tmi.twitch.tv 376 justinfan6493 :>
2023/07/11 09:14:12 :justinfan6493!justinfan6493@justinfan6493.tmi.twitch.tv JOIN #syluxdx
2023/07/11 09:14:12 :justinfan6493.tmi.twitch.tv 353 justinfan6493 = #syluxdx :justinfan6493
2023/07/11 09:14:12 :justinfan6493.tmi.twitch.tv 366 justinfan6493 #syluxdx :End of /NAMES list
```