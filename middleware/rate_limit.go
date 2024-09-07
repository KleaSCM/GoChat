package middleware

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu       sync.Mutex
	requests map[string]time.Time
	limit    int
	interval time.Duration
}

func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string]time.Time),
		limit:    limit,
		interval: interval,
	}
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rl.mu.Lock()
		defer rl.mu.Unlock()

		ip := r.RemoteAddr
		now := time.Now()
		if lastRequest, exists := rl.requests[ip]; exists {
			if now.Sub(lastRequest) < rl.interval {
				http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}
		}

		rl.requests[ip] = now
		next.ServeHTTP(w, r)
	})
}
