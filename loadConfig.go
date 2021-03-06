package aim_easy_tools

import (
	"encoding/json"
	"os"
)

//配置文件内容
//type configInfo struct {
//	RemoteEndpoint    string `json:"remote_endpoint"`    //服务器地址
//	LocalEndpoint     string `json:"local_endpoint"`     //本机地址
//	HttpTimeout       int64  `json:"http_timeout"`       //http请求超时时间
//	HeartbeatInterval int64  `json:"heartbeat_interval"` //心跳发送间隔
//
//}
//
//var (
//	ConfigInfo *configInfo
//	err        error
//)
//
//func init() {
//	c:=new(configInfo)
//	err = LoadConfig("config.json",c)
//	if err != nil {
//		log.Println(err.Error())
//		//time.Sleep(3000 * time.Millisecond)
//	}
//}

//读取配置文件
func LoadConfig(filePath string, config interface{}) error {
	cf, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer cf.Close()

	decoder := json.NewDecoder(cf)
	//config := configInfo{}
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	return nil
}
