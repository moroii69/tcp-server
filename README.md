# go-tcp-server

concurrent TCP server in go with per-client goroutines and channel-based message flow.

## features

- goroutine per client
- channel-based message handling
- raw TCP socket communication
- simple echo-style response

| server logs (`go run main.go`) | client (`ncat localhost 3000`) |
|-------------------------------|--------------------------------|
| <img src="https://i.postimg.cc/FFPrczbG/NVIDIA-Overlay-z9Eh6mj-H1e.png" width="320" /> | <img src="https://i.postimg.cc/W1ZTQ7BB/NVIDIA-Overlay-Dd-FVDNIrs-M.png" width="320" /> |

## run

```bash
go run main.go

```

## test

```
ncat localhost 3000
```
