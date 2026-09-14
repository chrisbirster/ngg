FROM node:24-alpine AS web
WORKDIR /src
RUN corepack enable
COPY package.json tsconfig.json vite.config.ts index.html ./
RUN pnpm install --no-frozen-lockfile
COPY src ./src
RUN pnpm build:web

FROM golang:1.26-alpine AS server
WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
COPY --from=web /src/internal/web/dist ./internal/web/dist
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/ngg ./cmd/ngg

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=server /out/ngg /ngg
ENV ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/ngg"]
