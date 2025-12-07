package common

import "time"

func GetTimeNow() time.Time {
	loc, _ := time.LoadLocation("Asia/Kuala_Lumpur")
	//set timezone,
	timeNow := time.Now().In(loc)
	return timeNow
}
