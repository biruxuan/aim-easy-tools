package websocket

import (
	"encoding/json"
	"testing"
)

func TestNewClient(t *testing.T) {
	//请假上报接口所需信息
	type leaveInfo struct {
		RequestID  int
		TerminalID string
		Time       int64
	}

	l := leaveInfo{
		RequestID:  12345,
		TerminalID: "mac_1234",
		Time:       2423453254325,
	}

	msg, err := json.Marshal(&l)
	if err != nil {
		t.Error(err.Error())
		return
	}

	c := NewClient()
	remote := "192.168.1.101:8083"

	err = c.Connect(remote, "/tempLeave")
	if err != nil {
		t.Error(err.Error())
		return
	}

	defer func() {
		if err := c.Close(); err != nil {
			t.Error(err.Error())
		}
	}()

	c.Write(msg)

}
