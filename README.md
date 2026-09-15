# NGG.GG

A creator-first browser game and video community inspired by the energy of classic web communities. NGG gives creators publishing, discovery, discussion, social features, an SDK, and an instant multiplayer arcade.

The first featured game is **XO Arena Football**.

## Run locally

```bash
corepack enable
pnpm install
pnpm dev
# separate terminal
pnpm dev:api
```

Open [http://localhost:5173/games/xo-arena-football](http://localhost:5173/games/xo-arena-football).

## Verify

```bash
pnpm verify
```

## Turso database

Production state is stored in the current Turso database service through its SQL-over-HTTP pipeline. The server creates its snapshot and append-only event tables at startup.

```bash
turso auth login
turso db create ngg --tursodb
export TURSO_DATABASE_URL="$(turso db show ngg --url)"
export TURSO_AUTH_TOKEN="$(turso db tokens create ngg)"
pnpm dev:api
```

Without `TURSO_DATABASE_URL`, the server deliberately uses ephemeral memory for local tests. A production Fly deployment should always set both secrets:

```bash
fly secrets set \
  TURSO_DATABASE_URL="$(turso db show ngg --url)" \
  TURSO_AUTH_TOKEN="$(turso db tokens create ngg)" \
  -a ngg
fly deploy --remote-only -a ngg
```

## Portal routes

- `/games/xo-arena-football` — existing featured game page
- `/discover` — games and video discovery
- `/submit` — creator publishing overview
- `/forums` — discussion categories
- `/arcade` — five instant multiplayer games
- `/account` — creator account signup

The complete API and milestone acceptance record live in [docs/community-platform.md](docs/community-platform.md).

## Branches

`feature/* → dev → release PR → main → tag`. See [docs/branching-and-releases.md](docs/branching-and-releases.md).
