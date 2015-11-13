# Client - Server communication API

Message_ID = first byte of message

## Client to Server
### Send_Message
Message_ID: 0
Sender:     (string)
Message:    (string)


## Server to Client
### Incoming_Message
Message_ID: 1
Sender:     (string)
Message:    (string)
