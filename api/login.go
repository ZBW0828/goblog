package api

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

//这个 API 接口处理函数实现了登录功能，接收用户名和密码，并返回对应的 JSON 数据作为响应。
func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	//接收用户名和密码，返回对应的json数据
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	LoginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, LoginRes)
}
