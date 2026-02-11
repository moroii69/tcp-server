# go-tcp-server

concurrent TCP server in go with per-client goroutines and channel-based message flow.

## features

- goroutine per client
- channel-based message handling
- raw TCP socket communication
- simple echo-style response

<img src="https://i.postimg.cc/FFPrczbG/NVIDIA-Overlay-z9Eh6mj-H1e.png" width="420" />

## run

```bash
go run main.go

```

## test

```
ncat localhost 3000
```
