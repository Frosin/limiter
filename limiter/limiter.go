package limiter

import "time"

type LimiterParams struct {
	TimeInterval time.Duration
	CountLimit   int
}

type Limiter interface {
	//Check returns true if request is possible
	Check(key string, timeStamp time.Time, params LimiterParams) (bool, error)
}
