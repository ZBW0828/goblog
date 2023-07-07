package service

import (
	"errors"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/utils"
)

//处理用户登录操作，包括验证用户名和密码是否正确、生成访问令牌以及返回用户信息和访问令牌。
func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "mszlu")
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid
	//生成token jwt技术进行生成
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil

}
