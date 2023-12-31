package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
)

//获取文章的归档信息，包括按照月份归档的文章列表和所有的分类列表。
func FindPostPigeonhole() models.PigeonholeRes {
	//查询所有文章，进行月份整理
	//查询所以分类
	posts, _ := dao.GetPostAll()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	categorys, _ := dao.GetAllCategory()
	return models.PigeonholeRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		categorys,
		pigeonholeMap,
	}

}
