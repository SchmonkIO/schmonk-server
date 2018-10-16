# Schmonk Game Server

JS Frontend can be found [here](https://github.com/SchmonkIO/schmonk-client-js).

## Requirements

- Go 1.9+

## Compile

- `go get -d`
- `go build`

## Example Config:

`server.conf`:
```
[SERVER]
IP = "123.456.789.000"
Port = 8080
TickRate = 20
Slots = 1000
CORS = false
Debug = true

[Game]
NameLength = 16
SlotsPerRoom = 10
```

## Actions

- SetUser: 
    - `{"action":"setUser","name":"..."}`
- GetRooms:
    - `{"action":"getRooms"}`
- CreateRoom:
    - `{"action":"createRoom","name":"...","pass":"","map":"","slots":4}`
- JoinRoom:
    - `{"action":"joinRoom","id":"...","pass":""}`
- LeaveRoom:
    - `{"action":"leaveRoom"}`
