package aim_easy_tools

import (
	"fmt"
	"strings"
)

func Byte2string(b []byte) string {
	str := fmt.Sprintf("%c", b[3:len(b)-1])
	body := strings.Replace(str, " ", "", -1)
	body = strings.Replace(body, "[", "", -1)
	body = strings.Replace(body, "]", "", -1)
	return body
}
