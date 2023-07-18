package views

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

//处理归档页面的请求，获取归档页面的模板，并调用服务层的函数获取归档页面的数据。然后，将数据写入响应。
func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole

	pigeonholeRes := service.FindPostPigeonhole()
	pigeonhole.WriteData(w, pigeonholeRes)
}
