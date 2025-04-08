package middleware

import (
	"SecureAPI/config"
	"log"
	"net/http"
	"strings"

	"github.com/rs/cors"
)

func MatchAllowedOrigin(origin string) bool {
	for _, allowed := range config.AppConfig.AllowedOrigins {
		if strings.HasPrefix(allowed, "*.") {
			suffix := strings.TrimPrefix(allowed, "*")
			if strings.HasSuffix(origin, suffix) {
				return true
			}
		} else if origin == allowed {
			return true
		}
	}
	return false
}

func StrictOriginEnforcer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		log.Printf("🔐 Strict check for origin: %s", origin)

		if origin != "" && !MatchAllowedOrigin(origin) {
			http.Error(w, "Forbidden: Origin not allowed", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetCORSHandler() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{}, // We'll use MatchAllowedOrigin instead
		AllowOriginFunc: func(origin string) bool {
			log.Printf("🌍 CORS check for origin: %s", origin)
			return MatchAllowedOrigin(origin)
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
}
