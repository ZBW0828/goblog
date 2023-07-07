package dao

import (
	"go-blog/models"
	"log"
)

//对数据库中分类信息的访问，可以根据分类的 ID 查询分类名称，或者获取所有的分类信息列表。
func GetCategoryNameById(cId int) string {
	row := DB.QueryRow("select name from blog_category where cid=?", cId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory 查询出错:", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory 取值出错:", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
