package views

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

//处理写作页面的请求，获取写作页面的模板，并调用服务层的函数获取写作页面需要的数据。然后，将数据写入响应。
func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}
