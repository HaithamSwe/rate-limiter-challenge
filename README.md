# rate-limiter-challenge
I requested a Golang coding challenge from ChatGPT, and it provided me with the following task:

## Problem: Build a Rate Limiter

### Description:
You need to design a rate limiter that allows a user to make a certain number of requests within a given time window. The rate limiter should support the following functionalities:

1. **Allow a request**: Check if a request is allowed based on the user’s request history and the defined rate limits.
2. **Record a request**: Record a request when it is made.

### Requirements:
1. Each user has their own rate limit.
2. The rate limit is defined by two parameters: `maxRequests` (the maximum number of requests allowed) and `timeWindow` (the time window in seconds).
3. The system should be able to handle multiple users independently.
4. The rate limiter should be efficient and scalable, capable of handling high throughput.

### Classes:
1. `RateLimiter` - This class will represent the rate limiter and manage the request limits for multiple users.

### Methods:
1. `AllowRequest(userID string) bool` - Check if a request is allowed for the given user. Return `true` if the request is allowed, `false` otherwise.
2. `RecordRequest(userID string)` - Record a request for the given user.

### Example:
```go
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
```

### Challenge:
1. Implement the `AllowRequest` method to check if a user’s request is within the allowed rate limit.
2. Implement the `RecordRequest` method to record a user’s request.
3. Ensure the system is efficient and can handle high throughput with multiple users.

Good luck!
