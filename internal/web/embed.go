package web

import (
	"embed"
	"io/fs"
	"mime"
	"net/http"
	"path"
	"strings"
)

//go:embed dist
var assets embed.FS

func Handler() http.Handler {
	sub, err := fs.Sub(assets, "dist")
	if err != nil { panic(err) }
	files := http.FileServer(http.FS(sub))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clean := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
		if clean == "." { clean = "index.html" }
		if f, err := sub.Open(clean); err == nil {
			_ = f.Close()
			if strings.HasPrefix(clean, "assets/") {
				w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			}
			if ext := path.Ext(clean); ext != "" {
				if value := mime.TypeByExtension(ext); value != "" { w.Header().Set("Content-Type", value) }
			}
			files.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Cache-Control", "no-cache")
		clone := r.Clone(r.Context())
		clone.URL.Path = "/"
		files.ServeHTTP(w, clone)
	})
}
