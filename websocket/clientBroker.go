package websocket

import (
	"github.com/gorilla/websocket"
	"net/url"
)

type wsClient interface {
	//连接服务器
	connect(remoteHost string, path string) error
	//关闭连接
	close() error
	//读取ws消息
	read() (int,[]byte, error)
	//发送ws消息
	write([]byte) error
}

type clientBroker struct {
	conn *websocket.Conn
	exit chan struct{}
}

func newClientBroker() *clientBroker {
	return &clientBroker{}
}

func (c *clientBroker) connect(remoteHost string, path string) error {
	u := url.URL{Scheme: "ws", Host: remoteHost, Path: path}

	var err error = nil
	c.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	return err
}

func (c *clientBroker) close() error {
	err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}

	return c.conn.Close()
}

func (c *clientBroker) read() (int ,[]byte, error) {
	mt, message, err := c.conn.ReadMessage()
	if err != nil {
		return mt, nil, err
	}
	return mt,message, nil
}

func (c *clientBroker) write(msg []byte) error {
	err := c.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return err
	}
	return nil
}
