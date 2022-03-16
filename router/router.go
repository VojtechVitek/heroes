package router

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"net/http"

	"github.com/VojtechVitek/heroes/rpc"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
)

const VERSION = "v0.0.1"

func Router(rpcServer *rpc.RPC) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.NoCache)            // Disable edge/browser cache, unless overridden explicitly.
	r.Use(middleware.Heartbeat("/ping")) // Defined before logger, since we don't want to flood the Kibana logs with /ping.
	r.Use(middleware.RealIP)             // Sets r.RemoteAddr from reverse-proxy headers (X-Forwarded-For, X-Real-IP).
	r.Use(middleware.RequestID)

	// Structured JSON logger for all HTTP requests.
	requestLogger := httplog.NewLogger("api", httplog.Options{
		JSON: VERSION != "development",
	})
	r.Use(httplog.Handler(requestLogger))

	if VERSION != "development" {
		// Recover from panics & print stack trace within a single JSON log line.
		// Must be defined after logger mw. This is required by chi.
		r.Use(middleware.Recoverer)
	} else {
		// Request/response logs with beautiful colored output.
		// Useful for development purposes only.
		//r.Use(middleware.Logger)
	}

	// Allow CORS from any domain. TODO: Proper CORS rules.
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			if r.Method == "OPTIONS" {
				w.WriteHeader(200)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	r.Get("/favicon.ico", favicon)
	r.Get("/robots.txt", robots)
	r.Get("/version", versionHandler)
	r.Get("/version.svg", versionSVG)

	r.Get("/maps/{mapName}", rpcServer.HandleMap)

	r.Handle("/*", rpc.NewAPIServer(rpcServer))

	return r
}

func renderErrorPage(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(500)
	fmt.Fprintf(w, "<h1>Error occurred</h1><br /><pre>%v</pre>\n", err)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

func robots(w http.ResponseWriter, r *http.Request) {
	// Disallow all robots. We don't want to be indexed by Google etc.
	fmt.Fprintf(w, "User-agent: *\nDisallow: /\n")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(VERSION))
}

func versionSVG(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	versionSVGImage.Execute(w, map[string]string{
		"env":     "aa-labs", // TODO: Inject an actual environment name.
		"version": VERSION,
		"color":   "#" + stringToRGB(VERSION),
	})
}

func stringToRGB(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash[0:6]
}

var versionSVGImage = template.Must(template.New("versionSVG").Parse(versionSVGImageTemplate))

const versionSVGImageTemplate = `
<svg xmlns="http://www.w3.org/2000/svg" width="400" height="20">
<linearGradient id="b" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="a">
<rect width="400" height="20" rx="3" fill="#fff"/></mask>
<g mask="url(#a)">
	<path fill="#555" d="M0 0h37v20H0z"/>
	<path fill="{{ .color }}" d="M37 0h400v20H37z"/>
	<path fill="url(#b)" d="M0 0h400v20H0z"/>
</g>
<g fill="#fff" text-anchor="left" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
	<text x="6" y="15" fill="#010101" fill-opacity=".3">{{ .env }}</text><text x="6" y="14">{{ .env }}</text>
	<text x="42" y="15" fill="#010101" fill-opacity=".3">{{ .version }}</text><text x="42" y="14">{{ .version }}</text>
</g></svg>
`
