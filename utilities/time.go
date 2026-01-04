package utilities

import "time"

func GetCurrentTimestampPlusSeconds(seconds int64) int64 {
	return time.Now().Add(time.Duration(seconds) * time.Second).Unix()
}

func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
