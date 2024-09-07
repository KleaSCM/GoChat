package middlewares

import (
	"net/http"
	"sync"
	"time"
)

var userMessageLimit = make(map[string]time.Time)
var mu sync.Mutex

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("X-Username")

		mu.Lock()
		defer mu.Unlock()

		if lastMessageTime, exists := userMessageLimit[username]; exists {
			if time.Since(lastMessageTime) < 2*time.Second {
				http.Error(w, "Too many messages, slow down", http.StatusTooManyRequests)
				return
			}
		}

		userMessageLimit[username] = time.Now()
		next.ServeHTTP(w, r)
	})
}
