package gitlab

import (
	"fmt"
	"sync"
)
import "github.com/Albert-Zhan/httpc"

type GLClient struct {
	baseUrl string
	token   string
}

var glClient *GLClient
var mu sync.Mutex

func NewGLClient(baseUrl string, token string) *GLClient {
	/*
		单例模式完成gitlab client的实例化
	*/
	if glClient == nil {
		mu.Lock()
		defer mu.Unlock()
		glClient = &GLClient{
			baseUrl: baseUrl,
			token:   token,
		}
	}
	return glClient
}

func (g *GLClient) ListUsers(username string) bool {
	url := g.baseUrl + fmt.Sprintf("/api/v4/users?username=%s", username)
	req := httpc.NewRequest(httpc.NewHttpClient())
	_, data, _ := req.SetUrl(url).SetHeader("PRIVATE-TOKEN", g.token).SetMethod("get").Send().End()
	fmt.Println(data)
	return true
}
