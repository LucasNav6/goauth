package utilities

import "time"

func GetCurrentTimestamp() time.Time {
	return time.Now()
}

func GetExpiryTimestamp(durationInSeconds int64) time.Time {
	return time.Now().Add(time.Duration(durationInSeconds) * time.Second)
}
