package service

import (
	"slogv2/src/main/entity"
	"slogv2/src/main/utils"
	"slogv2/src/main/utils/customError"
	"slogv2/src/main/vo"
	"strconv"
)

// CreateArticle 创建文章
// return status, err
func CreateArticle(article *entity.Article) (int, error) {
	if article.Title == "" || article.Content == "" {
		return customError.ARTICLE_CREATE_FAIL, customError.GetError(customError.ARTICLE_CREATE_FAIL, "标题或内容不能为空")
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
		return customError.ARTICLE_CREATE_FAIL, customError.GetError(customError.ARTICLE_CREATE_FAIL, err.Error())
	}
	if count == 0 {
		//文章总数为0时，Aid为0
		newArticle.Aid = 0
	} else {
		//查询最新的文章的Aid
		var _a entity.Article
		err = entity.Db.Last(&_a).Error
		if err != nil {
			return customError.ARTICLE_CREATE_FAIL, customError.GetError(customError.ARTICLE_CREATE_FAIL, err.Error())
		}
		//最新的文章的Aid加1
		newArticle.Aid = _a.Aid + 1
	}

	err = entity.Db.Create(&newArticle).Error
	if err != nil {
		return customError.ARTICLE_CREATE_FAIL, customError.GetError(customError.ARTICLE_CREATE_FAIL, err.Error())
	}
	return customError.SUCCESS, nil
}

// GetArticleByAid 根据aid获取文章
// return article, status, err
func GetArticleByAid(aid string) (entity.Article, int, error) {
	var article entity.Article
	err := entity.Db.Where("aid = ?", aid).Find(&article).Error
	if err != nil {
		return article, customError.ARTICLE_NOT_FOUND, customError.GetError(customError.ARTICLE_NOT_FOUND, err.Error())
	} else if article.Title == "" {
		return article, customError.ARTICLE_NOT_FOUND, customError.GetError(customError.ARTICLE_NOT_FOUND, "文章不存在或已被删除")
	}

	//TODO 点击量增加条件检测或单独编写点击量增加函数
	//		目前逻辑为每次get文章时点击量+1
	article.Clicks++
	err = entity.Db.Updates(&article).Error
	if err != nil {
		return article, customError.ARTICLE_UPDATE_FAIL, customError.GetError(customError.ARTICLE_UPDATE_FAIL, err.Error())
	}
	return article, customError.SUCCESS, nil
}

// UpdateArticle 更新文章
// return status, err
func UpdateArticle(article *entity.Article) (int, error) {
	if article.Title == "" || article.Content == "" {
		return customError.ARTICLE_CREATE_FAIL, customError.GetError(customError.ARTICLE_CREATE_FAIL, "标题或内容不能为空")
	}

	err := entity.Db.Where("aid =?", article.Aid).Updates(&article).Error
	if err != nil {
		return customError.ARTICLE_UPDATE_FAIL, customError.GetError(customError.ARTICLE_UPDATE_FAIL, err.Error())
	}
	return customError.SUCCESS, nil
}

// DeleteArticle 删除文章
// return status, err
func DeleteArticle(aid string) (int, error) {
	err := entity.Db.Where("aid =?", aid).Delete(&entity.Article{}).Error
	if err != nil {
		return customError.ARTICLE_DELETE_FAIL, customError.GetError(customError.ARTICLE_DELETE_FAIL, err.Error())
	}
	return customError.SUCCESS, nil
}

// UpdateArticleLikes 更新文章点赞数
// return status, err
func UpdateArticleLikes(aid string) (int, error) {
	var article entity.Article
	err := entity.Db.Where("aid =?", aid).Find(&article).Error
	if err != nil {
		return customError.ARTICLE_NOT_FOUND, customError.GetError(customError.ARTICLE_NOT_FOUND, err.Error())
	}
	if article.Title == "" {
		return customError.ARTICLE_NOT_FOUND, customError.GetError(customError.ARTICLE_NOT_FOUND, "文章不存在或已被删除")
	}
	article.Likes++
	err = entity.Db.Updates(&article).Error
	if err != nil {
		return customError.ARTICLE_UPDATE_FAIL, customError.GetError(customError.ARTICLE_UPDATE_FAIL, err.Error())
	}
	return customError.SUCCESS, nil
}

// GetArticleList 获取文章列表
// return articleList, count, status, err
func GetArticleList(page vo.Page) ([]entity.Article, int64, int, error) {
	var articles []entity.Article
	var total int64
	err := entity.Db.Model(&entity.Article{}).Count(&total).Error
	if err != nil {
		//log.Println(err)
		return nil, 0, customError.ARTICLE_LIST_FAIL, customError.GetError(customError.ARTICLE_LIST_FAIL, err.Error())
	}
	err = entity.Db.Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Find(&articles).Error
	if err != nil {
		//log.Println(err)
		return nil, 0, customError.ARTICLE_LIST_FAIL, customError.GetError(customError.ARTICLE_LIST_FAIL, err.Error())
	}
	if len(articles) == 0 {
		return nil, 0, customError.ARTICLE_LIST_FAIL, customError.GetError(customError.ARTICLE_LIST_FAIL, "没有更多文章了")
	}
	return articles, total, customError.SUCCESS, nil
}

// GetAchieveArticleList 获取归档文章列表
// param year 年份
// return AchieveArticle, status, err
func GetAchieveArticleList() ([]vo.AchieveArticle, int, error) {
	var achieveList []vo.AchieveArticle
	var years []int

	err := entity.Db.Table("article").Select("DISTINCT year(created_at) as year").Find(&years).Error
	if err != nil {
		return nil, customError.ARTICLE_LIST_FAIL, customError.GetError(customError.ARTICLE_LIST_FAIL, err.Error())
	}
	if len(years) == 0 {
		return nil, customError.ARTICLE_LIST_FAIL, customError.GetError(customError.ARTICLE_LIST_FAIL, "暂无文章")
	}

	for _, year := range years {
		var articleList []entity.Article
		err = entity.Db.Where("year(created_at) =?", year).Find(&articleList).Error
		if err != nil {
			return nil, customError.ARTICLE_LIST_FAIL, customError.GetError(customError.ARTICLE_LIST_FAIL, err.Error())
		}
		achieveList = append(achieveList, vo.AchieveArticle{
			Year:        year,
			ArticleList: articleList,
		})
	}
	return achieveList, customError.SUCCESS, nil
}

func CreateArticleWithCategory(articleWithCategory *vo.ArticleWithCategory) (int, error) {
	status, err := CreateArticle(&articleWithCategory.Article)
	if err != nil {
		return status, customError.GetError(status, err.Error())
	}

	status, err = AddCategoryList(articleWithCategory.Article, articleWithCategory.CategoryList)
	if err != nil {
		return status, customError.GetError(status, err.Error())
	}
	return customError.SUCCESS, nil
}

func UpdateArticleWithCategory(articleWithCategory *vo.ArticleWithCategory) (int, error) {
	status, err := UpdateArticle(&articleWithCategory.Article)
	if err != nil {
		return status, customError.GetError(status, err.Error())
	}

	status, err = DeleteRelationByAid(strconv.Itoa(articleWithCategory.Article.Aid))
	if err != nil {
		return status, customError.GetError(status, err.Error())
	}

	status, err = AddCategoryList(articleWithCategory.Article, articleWithCategory.CategoryList)
	if err != nil {
		return status, customError.GetError(status, err.Error())
	}

	return status, nil
}

func GetArticleWithCategory(aid string) (vo.ArticleWithCategory, int, error) {
	var articleWithCategory vo.ArticleWithCategory

	article, status, err := GetArticleByAid(aid)
	if err != nil {
		return articleWithCategory, status, customError.GetError(status, err.Error())
	}

	categoryList, status, err := GetCategoryListByAid(aid)
	if err != nil {
		return articleWithCategory, status, customError.GetError(status, err.Error())
	}

	articleWithCategory.Article = article
	articleWithCategory.CategoryList = categoryList

	return articleWithCategory, status, nil
}

func GetAllArticle() ([]entity.Article, int, int, error) {
	var articleList []entity.Article
	var total int64

	err := entity.Db.Model(&entity.Article{}).Count(&total).Error
	if err != nil {
		return nil, 0, customError.ARTICLE_LIST_FAIL, customError.GetError(customError.ARTICLE_LIST_FAIL, err.Error())
	}

	err = entity.Db.Find(&articleList).Error
	if err != nil {
		return nil, 0, customError.ARTICLE_LIST_FAIL, customError.GetError(customError.ARTICLE_LIST_FAIL, err.Error())
	}

	return articleList, int(total), customError.SUCCESS, nil
}

func GetArticleListInfo() (vo.Page, int, error) {
	var page vo.Page

	page.PageSize = utils.DefaultPageSize
	page.Page = 1

	err := entity.Db.Model(&entity.Article{}).Count(&page.Total).Error
	if err != nil {
		return page, customError.ARTICLE_LIST_FAIL, customError.GetError(customError.ARTICLE_LIST_FAIL, err.Error())
	}

	return page, customError.SUCCESS, nil
}
