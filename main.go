package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	maxRequests int
	timeWindow  time.Duration
	userRecords map[string][]time.Time
}

func NewRateLimiter(maxRequests int, timeWindow time.Duration) *RateLimiter {
	return &RateLimiter{
		maxRequests: maxRequests,
		timeWindow:  timeWindow,
		userRecords: make(map[string][]time.Time),
	}
}

func (r *RateLimiter) AllowRequest(userID string) bool {
	// Implement logic to check if the request is allowed
	return false
}

func (r *RateLimiter) RecordRequest(userID string) {
	// Implement logic to record a request
}

func main() {
	// Example Usage:
	rateLimiter := NewRateLimiter(5, time.Minute)

	userID := "user123"

	for i := 0; i < 10; i++ {
		if rateLimiter.AllowRequest(userID) {
			rateLimiter.RecordRequest(userID)
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request not allowed")
		}
		time.Sleep(10 * time.Second) // simulate time between requests
	}
}
