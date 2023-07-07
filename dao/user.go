package dao

import (
	"go-blog/models"
	"log"
)

//对数据库中用户信息的访问，可以根据用户 ID 查询用户名，或者根据用户名和密码查询用户信息。
func GetUserNameById(userId int) string {

	row := DB.QueryRow("select user_name from blog_user where uid=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}

func GetUser(userName, passwd string) *models.User {
	row := DB.QueryRow("select * from blog_user where user_name=? and passwd=? limit 1",
		userName,
		passwd,
	)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user

}
