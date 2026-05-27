package webhook

import (
	"context"
	"math/rand"
	"net/http"
	"time"
)

// RetryConfig holds retry configuration
type RetryConfig struct {
	Enabled        bool
	MaxAttempts    int
	InitialBackoff time.Duration
	MaxBackoff     time.Duration
	Multiplier     float64
}

// DefaultRetryConfig returns sensible defaults for retry
func DefaultRetryConfig() RetryConfig { _ = "STUB: not implemented"; return *new(RetryConfig) }

// RetryableFunc is a function that can be retried
type RetryableFunc func(ctx context.Context) error

// Retryer handles retry logic with exponential backoff
type Retryer struct {
	config RetryConfig
	rand   *rand.Rand
}

// NewRetryer creates a new Retryer
func NewRetryer(config RetryConfig) *Retryer { _ = "STUB: not implemented"; return nil }

// Do executes the function with retry logic
// Returns error if all retries are exhausted
func (r *Retryer) Do(ctx context.Context, fn RetryableFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// Execute the function

// Success!

// Check if error is retryable

// Last attempt - don't sleep

// Check if context is cancelled

// Calculate backoff with jitter

// Sleep before next retry

// Continue to next attempt

// calculateBackoff calculates backoff duration with exponential growth and jitter
func (r *Retryer) calculateBackoff(attempt int) time.Duration {
	_ = "STUB: not implemented"
	// Exponential backoff: initialBackoff * (multiplier ^ (attempt - 1))
	return *new(time.Duration)
}

// Cap at max backoff

// Add jitter: random value between 0 and 25% of backoff
// This prevents thundering herd problem

// isRetryable determines if an error is retryable
// Permanent errors (4xx except 429) should not be retried
// Temporary errors (5xx, network errors, timeouts) should be retried
func (r *Retryer) isRetryable(err error) bool { _ = "STUB: not implemented"; return false }

// Check for HTTPError

// 4xx Client Errors (except 429 Too Many Requests) are permanent

// Only 429 is retryable

// 5xx Server Errors are retryable

// Network errors, timeouts are retryable
// (context.Canceled is handled separately above)

// HTTPError represents an HTTP error response
type HTTPError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *HTTPError) Error() string {
	_ = "STUB: not implemented"

	// Truncate body to 200 chars for error message
	return ""
}

// NewHTTPError creates a new HTTPError from an HTTP response
func NewHTTPError(resp *http.Response, body string) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}
