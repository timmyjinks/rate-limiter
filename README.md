# rate-limiter

Implementations of common rate limiting algorithms from scratch in Go. Each algorithm is self-contained and runnable, making this a useful reference for understanding how rate limiting works under the hood.

## Algorithms

### Token Bucket
Tokens accumulate in a bucket at a fixed rate up to a maximum capacity. Each incoming request consumes one token. If the bucket is empty, the request is rejected. This allows short bursts of traffic while enforcing an average rate over time.

### Leaky Bucket
Requests enter a queue (the "bucket") and are processed at a fixed, constant rate like water leaking from a hole. Excess requests that overflow the bucket are dropped. Unlike token bucket, this algorithm smooths out bursts entirely.

### Fixed Window Counter
Requests are counted within discrete time windows (e.g., 100 requests per minute). The counter resets at the start of each window. Simple to implement, but susceptible to traffic spikes at window boundaries.

### Sliding Window Log
Tracks the exact timestamp of each request within a rolling time window. More accurate than fixed windows there are no boundary spikes but requires storing a log of recent request times.

## Project Structure

```
rate-limiter/
├── cmd/
│   ├── token_bucket_limiter.go       # Token bucket implementation
│   ├── leaky_bucket_limiter.go       # Leaky bucket implementation
│   ├── fixed_window_limiter.go       # Fixed window counter
│   └── sliding_window_limiter.go     # Sliding window log
└── go.mod
```

## Getting Started

**Prerequisites:** Go 1.18+

```bash
git clone https://github.com/timmyjinks/rate-limiter.git
cd rate-limiter
```

Run any algorithm directly:

```bash
go run ./cmd/token_bucket_limiter.go
go run ./cmd/leaky_bucket_limiter.go
go run ./cmd/fixed_window_limiter.go
go run ./cmd/sliding_window_limiter.go
```

## Algorithm Comparison

| Algorithm | Burst Handling | Memory Usage | Accuracy | Complexity |
|-----------|---------------|--------------|----------|------------|
| Token Bucket | ✅ Allows bursts | Low | Moderate | Low |
| Leaky Bucket | ❌ Smooths bursts | Low | High | Low |
| Fixed Window | ✅ Allows bursts | Very low | Low (boundary spikes) | Very low |
| Sliding Window Log | ✅ Allows bursts | High | Very high | Moderate |
