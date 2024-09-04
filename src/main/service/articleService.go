package service

import (
	"slogv2/src/main/entity"
	"slogv2/src/main/utils"
	"slogv2/src/main/utils/customError"
)

func CreateArticle(article *entity.Article) (int, error) {
	if article.Title == "" || article.Content == "" {
		return customError.ARTICLE_ADD_FAIL, customError.GetError(customError.ARTICLE_ADD_FAIL, "标题或内容不能为空")
	}

	var newArticle entity.Article
	newArticle.Title = article.Title
	newArticle.Content = article.Content
	newArticle.Uid = article.Uid
	newArticle.Likes = 0
	newArticle.Clicks = 0
	newArticle.AllowComment = article.AllowComment

	//摘要未填写时默认使用文章前100字
	//regex排除Markdown标签，仅保留文本
	if article.Abstract == "" {
		abs := utils.GetPureTextRegex(article.Content)
		if len(abs) > 100 {
			newArticle.Abstract = abs[:100]
		} else {
			newArticle.Abstract = abs
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
	//查询文章总数
	err := entity.Db.Find(&a).Count(&count).Error
	if err != nil {
		return customError.ARTICLE_ADD_FAIL, customError.GetError(customError.ARTICLE_ADD_FAIL, err.Error())
	}
	if count == 0 {
		//文章总数为0时，Aid为0
		newArticle.Aid = 0
	} else {
		//查询最新的文章的Aid
		var _a entity.Article
		err = entity.Db.Last(&_a).Error
		if err != nil {
			return customError.ARTICLE_ADD_FAIL, customError.GetError(customError.ARTICLE_ADD_FAIL, err.Error())
		}
		//最新的文章的Aid加1
		newArticle.Aid = _a.Aid + 1
	}

	err = entity.Db.Create(&newArticle).Error
	if err != nil {
		return customError.ARTICLE_ADD_FAIL, customError.GetError(customError.ARTICLE_ADD_FAIL, err.Error())
	}
	return customError.SUCCESS, nil
}

func GetArticleByAid(aid string) (entity.Article, int, error) {
	var article entity.Article
	err := entity.Db.Where("aid = ?", aid).Find(&article).Error
	if err != nil {
		return article, customError.ARTICLE_NOT_FOUND, customError.GetError(customError.ARTICLE_NOT_FOUND, err.Error())
	}
	article.Clicks++
	err = entity.Db.Updates(&article).Error
	if err != nil {
		return article, customError.ARTICLE_UPDATE_FAIL, customError.GetError(customError.ARTICLE_UPDATE_FAIL, err.Error())
	}
	return article, customError.SUCCESS, nil
}

func GetArticleList(page int, pageSize int) ([]entity.Article, int64, int, error) {
	var articles []entity.Article
	var total int64
	err := entity.Db.Model(&entity.Article{}).Count(&total).Error
	if err != nil {
		//log.Println(err)
		return nil, 0, customError.ARTICLE_LIST_FAIL, customError.GetError(customError.ARTICLE_LIST_FAIL, err.Error())
	}
	err = entity.Db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&articles).Error
	if err != nil {
		//log.Println(err)
		return nil, 0, customError.ARTICLE_LIST_FAIL, customError.GetError(customError.ARTICLE_LIST_FAIL, err.Error())
	}
	return articles, total, customError.SUCCESS, nil
}
