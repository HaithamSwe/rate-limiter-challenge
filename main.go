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
	if userRequestsTimes, exists := r.userRecords[userID]; !exists || len(userRequestsTimes) < r.maxRequests {
		return true
	}

	currentTime := time.Now()
	for i := 1; i <= r.maxRequests; i++ {
		if !r.userRecords[userID][len(r.userRecords[userID])-i].Before(currentTime.Add(-r.timeWindow)) {
			return false
		}
	}
	return true
}

func (r *RateLimiter) RecordRequest(userID string) {
	r.userRecords[userID] = append(r.userRecords[userID], time.Now())
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
