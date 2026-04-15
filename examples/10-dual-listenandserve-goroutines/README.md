Read this in: [🇧🇷 Português](README_pt.md) | **🇺🇸 English**

# Example 10: Dual ListenAndServe with Goroutines

This folder centralizes the dual-port example (`:8080` and `:3000`) and the blocking/shutdown strategies extracted from `versao-simples-1.go`.

Base endpoints used in all scenarios:
- `GET /api/v1/user` on `:8080`
- `POST /api/v1/user` on `:8080`
- `GET /api/v1/mock/user` on `:3000`
- `POST /api/v1/mock/user` on `:3000`

## Main Example

- [main.go](main.go): practical version with two servers, goroutines, signal handling, and graceful shutdown.

Run:

```bash
go run ./examples/10-dual-listenandserve-goroutines
```

## Scenarios

| # | Strategy | Path |
|---|---|---|
| 01 | Infinite block with `select {}` | [scenarios/01-select-block](scenarios/01-select-block) |
| 02 | Blocking channel (`<-done`) | [scenarios/02-blocking-channel](scenarios/02-blocking-channel) |
| 03 | `sync.WaitGroup` | [scenarios/03-waitgroup](scenarios/03-waitgroup) |
| 04 | Signal channel (`os/signal`) | [scenarios/04-signal-channel](scenarios/04-signal-channel) |
| 05 | `time.Sleep` | [scenarios/05-sleep](scenarios/05-sleep) |
| 06 | `context.WithCancel` | [scenarios/06-context-cancel](scenarios/06-context-cancel) |
| 07 | Infinite loop + sleep | [scenarios/07-infinite-loop](scenarios/07-infinite-loop) |
| 08 | Error channel group | [scenarios/08-error-channel-group](scenarios/08-error-channel-group) |
| 09 | Error channel + context | [scenarios/09-error-channel-group-context](scenarios/09-error-channel-group-context) |
| 10 | Mutex deadlock (didactic) | [scenarios/10-mutex-deadlock](scenarios/10-mutex-deadlock) |
| 11 | `runtime.Goexit()` | [scenarios/11-runtime-goexit](scenarios/11-runtime-goexit) |
| 12 | `sync.Cond` wait | [scenarios/12-sync-cond](scenarios/12-sync-cond) |
| 13 | `for range` on channel | [scenarios/13-channel-range](scenarios/13-channel-range) |
| 14 | Blocking `stdin` | [scenarios/14-stdin-block](scenarios/14-stdin-block) |
| 15 | One blocking server + one goroutine | [scenarios/15-one-blocking-server](scenarios/15-one-blocking-server) |
| 16 | Infinite ticker | [scenarios/16-infinite-ticker](scenarios/16-infinite-ticker) |
| 17 | `select` with multiple channels | [scenarios/17-select-multi-channels](scenarios/17-select-multi-channels) |

Run any scenario:

```bash
go run ./examples/10-dual-listenandserve-goroutines/scenarios/04-signal-channel
```
