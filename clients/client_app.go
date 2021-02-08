package clients

import "github.com/kataras/iris/v12"

type ClientApp struct {
	ID          string
	Name        string
	Secret      string
	EnabledRBAC bool
	Root        string
}

type ClientManager struct{}

func Get(c *iris.Context) *ClientManager {
	// 从实例中获取ClientManager实例
	return nil
}

// 中间件定义
func (c *ClientManager) Handler() func(*iris.Context) {
	return func(c *iris.Context) {
		//对Client进行验证，并获取Client的Secret
	}
}

func (c *ClientManager) All() map[string]interface{} {
	return nil
}

// 获取指定client_idApp
func (c *ClientManager) GetApp(client_id string) *ClientApp {
	return nil
}
