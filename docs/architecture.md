# Architecture

NGG is the portal; games remain independent applications.

```text
Browser → ngg.gg/games/:slug → Go API + embedded Solid SPA
                              ├─ manifest registry
                              ├─ reactions/comments/playlists/reports
                              └─ player adapter → local preview or external game URL
```

## Runtime

- Go 1.26 `net/http` serves `/api/v1/*` and the embedded SPA.
- Vite builds Solid 2 + Solid Router + StyleX into `internal/web/dist`.
- Game manifests are portable metadata. The player adapter decides whether a game is local, iframe-hosted, or launched externally.
- XO Arena stays owned by its game repository and deployment. NGG owns discovery and community behavior.

## Security boundary

Externally hosted games must use an allowlisted origin, sandboxed iframe permissions, explicit fullscreen policy, and postMessage schemas. NGG never grants a game access to portal session cookies.
