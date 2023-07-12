# Go Twicth Chatbot

|             | WebSocket clients               | IRC clients                   |
| ----------- | ------------------------------- | ----------------------------- |
| **SSL**     | wss://irc-ws.chat.twitch.tv:443 | irc://irc.chat.twitch.tv:6697 |
| **Non SSL** | ws://irc-ws.chat.twitch.tv:80   | irc://irc.chat.twitch.tv:6667 |


[Twitch chat docs](https://dev.twitch.tv/docs/irc/)  
[How to build websockets in Go](https://yalantis.com/blog/how-to-build-websockets-in-go/)  
[reader with timeout](https://gist.github.com/hongster/04660a20f2498fb7b680)

[golang console user interface](https://github.com/jroimartin/gocui)


# messages example
```
:tmi.twitch.tv 001 justinfan6493 :Welcome, GLHF!
:tmi.twitch.tv 002 justinfan6493 :Your host is tmi.twitch.tv
:tmi.twitch.tv 003 justinfan6493 :This server is rather new
:tmi.twitch.tv 004 justinfan6493 :-
:tmi.twitch.tv 375 justinfan6493 :-
:tmi.twitch.tv 372 justinfan6493 :You are in a maze of twisty passages, all alike.
:tmi.twitch.tv 376 justinfan6493 :>
:justinfan6493!justinfan6493@justinfan6493.tmi.twitch.tv JOIN #syluxdx
:justinfan6493.tmi.twitch.tv 353 justinfan6493 = #syluxdx :justinfan6493
:justinfan6493.tmi.twitch.tv 366 justinfan6493 #syluxdx :End of /NAMES list
:syluxdx!syluxdx@syluxdx.tmi.twitch.tv PRIVMSG #syluxdx :one message
:syluxdx!syluxdx@syluxdx.tmi.twitch.tv PRIVMSG #syluxdx :another message
PING :tmi.twitch.tv
```

# twitch example of message parsing
```javascript
function parseMessage(message) {

    let parsedMessage = {  // Contains the component parts.
        tags: null,
        source: null,
        command: null,
        parameters: null
    };

    // The start index. Increments as we parse the IRC message.

    let idx = 0; 

    // The raw components of the IRC message.

    let rawTagsComponent = null;
    let rawSourceComponent = null; 
    let rawCommandComponent = null;
    let rawParametersComponent = null;

    // If the message includes tags, get the tags component of the IRC message.

    if (message[idx] === '@') {  // The message includes tags.
        let endIdx = message.indexOf(' ');
        rawTagsComponent = message.slice(1, endIdx);
        idx = endIdx + 1; // Should now point to source colon (:).
    }

    // Get the source component (nick and host) of the IRC message.
    // The idx should point to the source part; otherwise, it's a PING command.

    if (message[idx] === ':') {
        idx += 1;
        let endIdx = message.indexOf(' ', idx);
        rawSourceComponent = message.slice(idx, endIdx);
        idx = endIdx + 1;  // Should point to the command part of the message.
    }

    // Get the command component of the IRC message.

    let endIdx = message.indexOf(':', idx);  // Looking for the parameters part of the message.
    if (-1 == endIdx) {                      // But not all messages include the parameters part.
        endIdx = message.length;                 
    }

    rawCommandComponent = message.slice(idx, endIdx).trim();

    // Get the parameters component of the IRC message.

    if (endIdx != message.length) {  // Check if the IRC message contains a parameters component.
        idx = endIdx + 1;            // Should point to the parameters part of the message.
        rawParametersComponent = message.slice(idx);
    }

    // Parse the command component of the IRC message.

    parsedMessage.command = parseCommand(rawCommandComponent);

    // Only parse the rest of the components if it's a command
    // we care about; we ignore some messages.

    if (null == parsedMessage.command) {  // Is null if it's a message we don't care about.
        return null; 
    }
    else {
        if (null != rawTagsComponent) {  // The IRC message contains tags.
            parsedMessage.tags = parseTags(rawTagsComponent);
        }

        parsedMessage.source = parseSource(rawSourceComponent);

        parsedMessage.parameters = rawParametersComponent;
        if (rawParametersComponent && rawParametersComponent[0] === '!') {  
            // The user entered a bot command in the chat window.            
            parsedMessage.command = parseParameters(rawParametersComponent, parsedMessage.command);
        }
    }

    return parsedMessage;
}
```