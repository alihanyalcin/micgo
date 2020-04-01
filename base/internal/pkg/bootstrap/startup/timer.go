package startup

import "time"

// Timer contains references to dependencies required by the startup timer implementation.
type Timer struct {
	startTime time.Time
	duration  time.Duration
	interval  time.Duration
}

// NewStartUpTimer is a factory method that returns an initialized Timer receiver struct.
func NewStartUpTimer(retryIntervalInSeconds, maxWaitInSeconds int) Timer {
	return Timer{
		startTime: time.Now(),
		duration:  time.Second * time.Duration(maxWaitInSeconds),
		interval:  time.Second * time.Duration(retryIntervalInSeconds),
	}
}

// SinceAsString returns the time since the timer was created as a string.
func (t Timer) SinceAsString() string {
	return time.Since(t.startTime).String()
}

// HasNotElapsed returns whether or not the duration specified during construction has elapsed.
func (t Timer) HasNotElapsed() bool {
	return time.Now().Before(t.startTime.Add(t.duration))
}

// SleepForInterval pauses execution for the interval specified during construction.
func (t Timer) SleepForInterval() {
	time.Sleep(t.interval)
}

//	Update the wait/interval for the timer,
func (t Timer) UpdateTimer(maxWaitInSeconds int, retryIntervalInSeconds int) {
	if maxWaitInSeconds > 0 {
		t.duration = time.Second * time.Duration(maxWaitInSeconds)
	}
	if retryIntervalInSeconds > 0 {
		t.interval = time.Second * time.Duration(retryIntervalInSeconds)
	}
}
