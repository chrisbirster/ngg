# NGG.GG

A creator-first browser game portal inspired by the energy of classic web-game communities. NGG gives independently deployed HTML5/WASM games a shared discovery, presentation, and social layer.

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

## Branches

`feature/* → dev → release PR → main → tag`. See [docs/branching-and-releases.md](docs/branching-and-releases.md).
