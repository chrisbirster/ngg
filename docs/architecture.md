# Architecture

NGG is the portal and the initial home of XO Arena Football.

```text
Browser → ngg.gg/games/:slug → Go API + embedded Solid SPA
                              ├─ manifest registry
                              ├─ reactions/comments/playlists/reports
                              └─ /play/:slug → game runtime
```

## Runtime

- Go 1.26 `net/http` serves `/api/v1/*` and the embedded SPA.
- Vite builds Solid 2 + Solid Router + StyleX into `internal/web/dist`.
- Game manifests use relative NGG URLs by default, so a game never depends on an unavailable third-party or retired domain.
- XO Arena is branded and launched entirely through NGG. Its canonical game path is `/play/xo-arena-football`.
- The manifest boundary still permits a future independently deployed runtime, but only after NGG controls and verifies its reachable origin.

## Security boundary

A future externally hosted game must use an NGG-controlled allowlisted origin, sandboxed iframe permissions, an explicit fullscreen policy, and versioned `postMessage` schemas. NGG never grants a game access to portal session cookies.
