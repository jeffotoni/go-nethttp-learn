Leia em: **🇧🇷 Português** | [🇺🇸 English](README.md)

# Exemplo 10: Dual ListenAndServe com Goroutines

Esta pasta centraliza o exemplo de duas portas (`:8080` e `:3000`) e as estratégias de bloqueio/parada extraídas do `versao-simples-1.go`.

Endpoints base usados em todos os cenários:
- `GET /api/v1/user` na `:8080`
- `POST /api/v1/user` na `:8080`
- `GET /api/v1/mock/user` na `:3000`
- `POST /api/v1/mock/user` na `:3000`

## Exemplo Principal

- [main.go](main.go): versão prática com dois servidores, goroutines, captura de sinais e graceful shutdown.

Executar:

```bash
go run ./examples/10-dual-listenandserve-goroutines
```

## Cenários

| # | Estratégia | Caminho |
|---|---|---|
| 01 | Bloqueio infinito com `select {}` | [scenarios/01-select-block](scenarios/01-select-block) |
| 02 | Channel bloqueante (`<-done`) | [scenarios/02-blocking-channel](scenarios/02-blocking-channel) |
| 03 | `sync.WaitGroup` | [scenarios/03-waitgroup](scenarios/03-waitgroup) |
| 04 | Canal de sinal (`os/signal`) | [scenarios/04-signal-channel](scenarios/04-signal-channel) |
| 05 | `time.Sleep` | [scenarios/05-sleep](scenarios/05-sleep) |
| 06 | `context.WithCancel` | [scenarios/06-context-cancel](scenarios/06-context-cancel) |
| 07 | Loop infinito + sleep | [scenarios/07-infinite-loop](scenarios/07-infinite-loop) |
| 08 | Grupo por canal de erro | [scenarios/08-error-channel-group](scenarios/08-error-channel-group) |
| 09 | Canal de erro + context | [scenarios/09-error-channel-group-context](scenarios/09-error-channel-group-context) |
| 10 | Deadlock com mutex (didático) | [scenarios/10-mutex-deadlock](scenarios/10-mutex-deadlock) |
| 11 | `runtime.Goexit()` | [scenarios/11-runtime-goexit](scenarios/11-runtime-goexit) |
| 12 | Espera com `sync.Cond` | [scenarios/12-sync-cond](scenarios/12-sync-cond) |
| 13 | `for range` em channel | [scenarios/13-channel-range](scenarios/13-channel-range) |
| 14 | `stdin` bloqueante | [scenarios/14-stdin-block](scenarios/14-stdin-block) |
| 15 | Um servidor bloqueante + um em goroutine | [scenarios/15-one-blocking-server](scenarios/15-one-blocking-server) |
| 16 | Ticker infinito | [scenarios/16-infinite-ticker](scenarios/16-infinite-ticker) |
| 17 | `select` com múltiplos canais | [scenarios/17-select-multi-channels](scenarios/17-select-multi-channels) |

Executar qualquer cenário:

```bash
go run ./examples/10-dual-listenandserve-goroutines/scenarios/04-signal-channel
```
