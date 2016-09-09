package main

import (
	"testing"
)

func TestRateLimiter(t *testing.T) {
	t.Log("Exceeding rate limit of 100/60...")

	r := NewRateLimiter()
	var allowed bool
	for i := 0; i < 101; i++ {
		allowed = r.allowed("some-key")
	}

	if allowed {
		t.Fatalf("Expected to be rate limited, but was not")
	}
}
