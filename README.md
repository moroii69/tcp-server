# go-tcp-server

concurrent TCP server in go with per-client goroutines and channel-based message flow.

## features

- goroutine per client
- channel-based message handling
- raw TCP socket communication
- simple echo-style response

## run

```bash
go run main.go

```

## test

```
ncat localhost 3000
```