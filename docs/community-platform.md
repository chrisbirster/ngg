# NGG community platform

This document is the acceptance record for the ten creator-community milestones. The implementation is a production-shaped vertical slice: its state and audit events survive Fly restarts in Turso, its APIs are bounded, and external media systems have explicit handoff points.

## Architecture

The Go server owns identities, publishing workflow, community state, moderation, SDK state, and authoritative arcade-room versions. Solid renders the portal and game pages. Turso stores one canonical JSON snapshot plus an append-only event for every mutation. This keeps the initial schema small while preserving a clean future migration path to normalized tables.

Turso tables created at startup:

```sql
CREATE TABLE platform_snapshots (
  key TEXT PRIMARY KEY,
  payload TEXT NOT NULL,
  updated_at TEXT NOT NULL
);

CREATE TABLE platform_events (
  id TEXT PRIMARY KEY,
  event_type TEXT NOT NULL,
  payload TEXT NOT NULL,
  created_at TEXT NOT NULL
);
```

The server uses `POST /v2/pipeline` directly with bearer authentication. No deprecated libSQL client, CGO, local replica, or filesystem database is involved.

## Milestone acceptance

| # | Capability | Implemented surface |
|---|---|---|
| 1 | Accounts and profiles | Account creation, unique email/handle, public profiles, opaque hashed sessions |
| 2 | Creator projects | Game/video drafts, metadata, tags, owner/collaborator roles, submission |
| 3 | Upload pipeline | Upload tickets, immutable object keys, bounded metadata, completion state |
| 4 | Game publishing | Completed-upload validation, immutable releases, `/play/{slug}` player URL boundary |
| 5 | Video publishing | Video release begins `processing`, `/watch/{slug}` boundary, transcoder status callback |
| 6 | Social graph | Polymorphic comments, 1–5 ratings, favorites, plus existing likes/playlists/reports |
| 7 | Forums | Categories, threads, replies, locked-thread guard, `@handle` notifications |
| 8 | Trust and safety | Generic reports, open moderation queue, resolution actions, Turso audit events |
| 9 | Game SDK | Idempotent achievements, top-100 leaderboards, optimistic/versioned cloud saves |
| 10 | Instant arcade | Five game manifests, rooms, joins, authoritative moves, optimistic room versions |

## Important media boundary

Turso stores metadata and platform state, not game ZIPs or video bytes. Upload ticket records contain the immutable object key that an R2/S3 signer will authorize. The current local ticket URL is a development seam; production must connect it to an object-store signer. Likewise, video releases deliberately remain `processing` until a transcoder webhook marks a renditon ready.

Uploaded games must ultimately run on a separate origin such as `play.ngg.gg`, inside an iframe with a narrow `sandbox` policy. Never execute creator HTML on the authenticated `ngg.gg` origin.

## API groups

- `/api/v1/accounts`, `/profiles`, `/sessions`
- `/api/v1/projects`, `/uploads`, `/releases`
- `/api/v1/content/{type}/{id}`
- `/api/v1/forums`, `/notifications`
- `/api/v1/reports`, `/moderation`
- `/api/v1/sdk/games/{gameID}`
- `/api/v1/arcade/games`, `/arcade/rooms`
- `/api/v1/healthz`, `/api/v1/readyz`

## Operational checks

```bash
curl -fsS https://ngg.fly.dev/api/v1/healthz
curl -fsS https://ngg.fly.dev/api/v1/readyz
curl -fsS https://ngg.fly.dev/api/v1/arcade/games
curl -fsS https://ngg.fly.dev/api/v1/forums/categories
```

`healthz` proves the process is alive. `readyz` includes a live Turso query when persistence is configured.
