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
	userRequestsTimes, userIdExists := r.userRecords[userID]
	if  !userIdExists || len(userRequestsTimes) == 0 {
		return true
	}

	return userRequestsTimes[len(userRequestsTimes)-1].Before(time.Now().Add(-r.timeWindow)) 
}

func (r *RateLimiter) RecordRequest(userID string) {
	if _, userIdExists := r.userRecords[userID]; !userIdExists {
		r.userRecords[userID] = []time.Time{}
	}
	
	if len(r.userRecords[userID]) == r.maxRequests {
		r.userRecords[userID] = r.userRecords[userID][1:]
	}
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
