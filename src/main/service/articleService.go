package service

import (
	"log"
	"slogv2/src/main/entity"
	"slogv2/src/main/utils"
	"slogv2/src/main/utils/customError"
)

func CreateArticle(article *entity.Article) int {
	if article.Title == "" || article.Content == "" {
		return customError.FAIL
	}

	var newArticle entity.Article
	newArticle.Title = article.Title
	newArticle.Content = article.Content
	newArticle.Uid = article.Uid
	newArticle.Likes = 0
	newArticle.Clicks = 0
	newArticle.AllowComment = article.AllowComment

	//摘要未填写时默认使用文章前100字
	if article.Abstract == "" {
		if len(article.Content) > 100 {
			newArticle.Abstract = article.Content[:100]
		} else {
			newArticle.Abstract = article.Content
		}
	} else {
		newArticle.Abstract = article.Abstract
	}

	//TODO 实现文件管理后上传图片至OSS后生成头图链接并传至前端暂存，在保存文章时再传回后端
	//头图链接未使用时使用默认图片
	if article.Image == "" {
		newArticle.Image = utils.TestDefaultImg
	} else {
		newArticle.Image = article.Image
	}

	//手动实现主键Aid自增
	var a entity.Article
	var count int64
	err := entity.Db.Find(&a).Count(&count).Error
	if err != nil {
		log.Println(err)
		return customError.FAIL
	}
	if count == 0 {
		newArticle.Aid = 0
	} else {
		err = entity.Db.Last(&a).Error
		if err != nil {
			log.Println(err)
			return customError.FAIL
		}
		newArticle.Aid = a.Aid + 1
	}

	err = entity.Db.Create(&newArticle).Error
	if err != nil {
		log.Println(err)
		return customError.FAIL
	}
	return customError.SUCCESS
}

func GetArticleList(page int, pageSize int) ([]entity.Article, int64, int) {
	var articles []entity.Article
	var total int64
	err := entity.Db.Model(&entity.Article{}).Count(&total).Error
	if err != nil {
		log.Println(err)
		return nil, 0, customError.FAIL
	}
	err = entity.Db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&articles).Error
	if err != nil {
		log.Println(err)
		return nil, 0, customError.FAIL
	}
	return articles, total, customError.SUCCESS
}
