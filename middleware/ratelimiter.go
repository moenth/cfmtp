package middleware

import (
	"log"
	"time"

	"github.com/kataras/iris"
)

// RateLimiter rate limits incoming API calls.
// A distributed implementation could use something like redis
// to share limiter state, but we'll keep it simple here.
// We also want to clean out old sessions periodically
// to avoid memory leaks.
type RateLimiter struct {
	Sessions map[string]*ratelimit
}

// Serve handles the request.
func (r RateLimiter) Serve(c *iris.Context) {
	key := c.RequestHeader("API_KEY")

	// Carry on if the allowance limit has not been reached.
	if r.allowed(key) {
		c.Next()
		return
	}

	// Returns 429 too many requests if allowance limit is exceeded
	log.Printf("Api key exceeded allowance: %s", key)
	c.JSON(429, "API call allowance exceeded. Try again later.")
}

// Clean removes old sessions.
func (r RateLimiter) Clean() {
	for key, l := range r.Sessions {
		if time.Since(l.lastCheck).Seconds() >= l.limit {
			delete(r.Sessions, key)
		}
	}
}

// Allowed checks if the key has any allowances left.
// A single api key is limited to 100 api calls every 60 seconds.
func (r RateLimiter) allowed(key string) bool {

	// For demonstration purposes an empty key can be used for load testing.
	if key == "" {
		return true
	}

	limit, ok := r.Sessions[key]
	if !ok {
		limit = newLimit(100, 60)
		r.Sessions[key] = limit
	}

	return limit.check()
}

// NewRateLimiter creates a new RateLimiter middleware.
func NewRateLimiter() RateLimiter {
	return RateLimiter{
		Sessions: map[string]*ratelimit{},
	}
}

// Ratelimit contains information about a rate limited entity.
type ratelimit struct {
	rate      int
	allowance int
	limit     float64
	lastCheck time.Time
}

// Check asserts if the limit has any allowances left.
func (l *ratelimit) check() bool {

	// Reset the allowance if a second has elapsed.
	if time.Since(l.lastCheck).Seconds() >= l.limit {
		l.allowance = l.rate
		l.lastCheck = time.Now()
	}

	l.allowance--
	return l.allowance >= 0
}

// newLimit creates a new rate limit.
func newLimit(rate int, limit float64) *ratelimit {
	return &ratelimit{
		rate:      rate,
		allowance: rate,
		limit:     limit,
		lastCheck: time.Now(),
	}
}
