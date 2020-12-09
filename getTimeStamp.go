package aim_easy_tools

import "time"

//返回ms
func GetTimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}
