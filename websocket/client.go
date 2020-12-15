package websocket

type client struct {
	ws wsClient
}

func NewClient() *client {
	return &client{
		ws: newClientBroker(),
	}
}

func (c *client) Connect(remoteHost string, path string) error {
	return c.ws.connect(remoteHost, path)
}

func (c *client) Close() error {

	return c.ws.close()
}

func (c *client) Read() (int,[]byte, error) {
	return c.ws.read()
}

func (c *client) Write(msg []byte) error {
	return c.ws.write(msg)
}
