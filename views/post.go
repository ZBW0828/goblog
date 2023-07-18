package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

//处理文章详情页面的请求，从请求中获取文章的 ID，并调用服务层的函数获取文章的详细信息。然后，将详细信息写入响应。
func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	//获取路径参数
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	//7.html
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别次请求路径"))
		return
	}
	postRes, err := service.GetPostDetail(pid)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)
}
