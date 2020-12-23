package convert

import "time"

// ParseTimeRange 會將 UNIX timeStamp(sec) 轉成 time.Time
func ParseTimeRange(beginTimeStamp, endTimeStamp int64) (time.Time, time.Time) {
	var beginDate, endDate time.Time
	if beginTimeStamp != 0 {
		beginDate = time.Unix(beginTimeStamp, 0)
	}
	if endTimeStamp != 0 {
		endDate = time.Unix(endTimeStamp, 0)
	}

	return beginDate, endDate
}
