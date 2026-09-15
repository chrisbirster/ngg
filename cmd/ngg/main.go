package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chrisbirster/ngg/internal/platform"
	"github.com/chrisbirster/ngg/internal/web"
)

func main() {
	ctx := context.Background()
	store := platform.NewStore()
	if databaseURL:=os.Getenv("TURSO_DATABASE_URL"); databaseURL!="" {
		backend,err:=platform.NewTursoBackend(databaseURL,os.Getenv("TURSO_AUTH_TOKEN"),nil)
		if err!=nil { slog.Error("configure Turso", "error",err); os.Exit(1) }
		if err:=backend.Migrate(ctx);err!=nil { slog.Error("migrate Turso", "error",err); os.Exit(1) }
		store,err=platform.NewPersistentStore(ctx,backend)
		if err!=nil { slog.Error("load Turso state", "error",err); os.Exit(1) }
		slog.Info("Turso persistence enabled")
	} else { slog.Warn("TURSO_DATABASE_URL is unset; using ephemeral development storage") }
	mux := http.NewServeMux()
	mux.Handle("/api/", platform.NewHandler(store))
	mux.Handle("/", web.Handler())

	addr := os.Getenv("ADDR")
	if addr == "" { addr = ":8080" }
	server := &http.Server{Addr: addr, Handler: securityHeaders(mux), ReadHeaderTimeout: 5*time.Second, IdleTimeout: 60*time.Second}
	stop, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	go func() {
		<-stop.Done()
		ctx, done := context.WithTimeout(context.Background(), 10*time.Second); defer done()
		_ = server.Shutdown(ctx)
	}()
	slog.Info("ngg listening", "addr", addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err,http.ErrServerClosed) { slog.Error("server stopped", "error", err); os.Exit(1) }
}
func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		w.Header().Set("X-Content-Type-Options","nosniff")
		w.Header().Set("Referrer-Policy","strict-origin-when-cross-origin")
		w.Header().Set("Permissions-Policy","camera=(), microphone=(), geolocation=()")
		next.ServeHTTP(w,r)
	})
}
