package aim_easy_tools

import "testing"

func TestLoadConfig(t *testing.T) {
	//配置文件内容
	type configInfo struct {
		RemoteEndpoint    string `json:"remote_endpoint"`
		LocalEndpoint     string `json:"local_endpoint"`
		HttpTimeout       int64  `json:"http_timeout"`
		HeartbeatInterval int64  `json:"heartbeat_interval"`
	}

	c:=new(configInfo)
	err:=LoadConfig("config.json",c)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(c)
}
