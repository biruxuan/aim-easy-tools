package aim_easy_tools

import "time"

func GetTimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}
