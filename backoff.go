package twitter

import "time"

const (
	maxNetDelay      = 16 * time.Second
	netDelayInc      = 250 * time.Millisecond
	maxHTTPDelay     = 320 * time.Second
	httpInitDelay    = 5 * time.Second
	http420InitDelay = time.Minute
)

// Backoff is the interface for stream retries. When an error is encountered,
// the stream's StreamErrFn will be called (if it exists) with the current
// Backoff and the sepcific error.
type Backoff interface {
	// NextWait returns the duration that the stream will wait before
	// attempting to reconnect.
	NextWait() time.Duration
	// Waited returns the total duration that has been waited (not including
	// the NextWait duration) since a successful stream was established.
	Waited() time.Duration
	// Retries returns the total number of retries since a successful stream
	// was established.
	Retries() int
}

type backoff struct {
	netDelay  time.Duration
	httpDelay time.Duration
	waited    time.Duration
	retries   int
}

func (b *backoff) NextWait() time.Duration {
	if b.netDelay > b.httpDelay {
		return b.netDelay
	}
	return b.httpDelay
}

func (b *backoff) Waited() time.Duration {
	return b.waited
}

func (b *backoff) Retries() int {
	return b.retries
}

func (b *backoff) reset() {
	b.netDelay = 0
	b.httpDelay = 0
	b.waited = 0
	b.retries = 0
}

func (b *backoff) wait() time.Duration {
	var wait time.Duration
	if b.netDelay > b.httpDelay {
		wait = b.netDelay
	} else {
		wait = b.httpDelay
	}
	b.waited += wait
	b.retries++
	return wait
}

func (b *backoff) incNetDelay() {
	b.netDelay += netDelayInc
	if b.netDelay > maxNetDelay {
		b.netDelay = maxNetDelay
	}
}

func (b *backoff) incHTTPDelay(is420 bool) {
	if b.httpDelay <= 0 {
		if is420 {
			b.httpDelay = http420InitDelay
			return
		}
		b.httpDelay = httpInitDelay
		return
	}
	b.httpDelay *= 2
	if b.httpDelay > maxHTTPDelay {
		b.httpDelay = maxHTTPDelay
	}
}
