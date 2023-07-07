package views

import (
	"go-blog/common"
	"go-blog/config"
	"net/http"
)

//处理登录页面的请求，获取登录页面的模板，并将配置中的 Viewer 数据写入响应。
func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)
}
