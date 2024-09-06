package service

import (
	"slogv2/src/main/entity"
	"slogv2/src/main/utils/customError"
	"slogv2/src/main/vo"
)

func CreateComment(comment *entity.Comment) (int, error) {
	if comment.Content == "" || comment.Uname == "" {
		return customError.COMMENT_CREATE_FAIL, customError.GetError(customError.COMMENT_CREATE_FAIL, "评论内容不能为空")
	}
	comment.Likes = 0
	comment.Dislikes = 0

	//根据配置判断评论是否需要审核
	options, status, err := GetOptionByKey("comment_need_audit")
	if status != customError.SUCCESS {
		return status, customError.GetError(customError.COMMENT_CREATE_FAIL, customError.GetMsg(status))
	}
	if options.Value == "true" {
		comment.IsActive = false
	} else {
		comment.IsActive = true
	}

	var count int64
	err = entity.Db.Find(&entity.Comment{}).Count(&count).Error
	if err != nil {
		return customError.COMMENT_CREATE_FAIL, customError.GetError(customError.COMMENT_CREATE_FAIL, err.Error())
	}

	if count == 0 {
		comment.Cid = 0
	} else {
		var _comment entity.Comment
		err := entity.Db.Last(&_comment).Error
		if err != nil {
			return customError.COMMENT_CREATE_FAIL, customError.GetError(customError.COMMENT_CREATE_FAIL, err.Error())
		}
		comment.Cid = _comment.Cid + 1
	}

	err = entity.Db.Create(comment).Error
	if err != nil {
		return customError.COMMENT_CREATE_FAIL, customError.GetError(customError.COMMENT_CREATE_FAIL, err.Error())
	}

	return customError.SUCCESS, nil
}

// GetCommentListByAid 获取评论列表
// 通过Aid获取评论列表
func GetCommentListByAid(aid string) ([]entity.Comment, int, error) {
	var commentList []entity.Comment
	err := entity.Db.Where("aid =?", aid).Find(&commentList).Error
	if err != nil {
		return commentList, customError.COMMENT_LIST_FAIL, customError.GetError(customError.COMMENT_LIST_FAIL, err.Error())
	}

	return commentList, customError.SUCCESS, nil
}

// GetCommentTreeByAid 获取评论树
// 通过Aid获取评论树
func GetCommentTreeByAid(aid string) (vo.CommentNode, int, error) {
	commentList, status, err := GetCommentListByAid(aid)

	var commentTree vo.CommentNode
	commentTree = vo.BuildCommentTree(commentList)

	return commentTree, status, err
}

func GetCommentList() ([]entity.Comment, int, error) {
	var commentList []entity.Comment
	err := entity.Db.Find(&commentList).Error
	if err != nil {
		return commentList, customError.COMMENT_LIST_FAIL, customError.GetError(customError.COMMENT_LIST_FAIL, err.Error())
	}

	return commentList, customError.SUCCESS, nil
}

func GetPendingCommentList() ([]entity.Comment, int, error) {
	var commentList []entity.Comment
	err := entity.Db.Where("is_active =?", false).Find(&commentList).Error
	if err != nil {
		return commentList, customError.COMMENT_LIST_FAIL, customError.GetError(customError.COMMENT_LIST_FAIL, err.Error())
	}

	return commentList, customError.SUCCESS, nil
}

func GetCommentByCid(cid string) (entity.Comment, int, error) {
	var comment entity.Comment
	err := entity.Db.Where("cid =?", cid).First(&comment).Error
	if err != nil {
		return comment, customError.COMMENT_NOT_FOUND, customError.GetError(customError.COMMENT_NOT_FOUND, err.Error())
	}

	return comment, customError.SUCCESS, nil
}

func UpdateComment(comment *entity.Comment) (int, error) {
	if comment.Content == "" || comment.Uname == "" {
		return customError.COMMENT_UPDATE_FAIL, customError.GetError(customError.COMMENT_UPDATE_FAIL, "评论内容不能为空")
	}

	err := entity.Db.Where("cid =?", comment.Cid).Updates(comment).Error
	if err != nil {
		return customError.COMMENT_UPDATE_FAIL, customError.GetError(customError.COMMENT_UPDATE_FAIL, err.Error())
	}

	return customError.SUCCESS, nil
}

func DeleteComment(cid string) (int, error) {
	err := entity.Db.Where("cid =?", cid).Delete(&entity.Comment{}).Error
	if err != nil {
		return customError.COMMENT_DELETE_FAIL, customError.GetError(customError.COMMENT_DELETE_FAIL, err.Error())
	}

	return customError.SUCCESS, nil
}

func GetCommentCountByAid(aid string) (int64, int, error) {
	var count int64
	err := entity.Db.Where("aid =?", aid).Count(&count).Error
	if err != nil {
		return count, customError.COMMENT_LIST_FAIL, customError.GetError(customError.COMMENT_LIST_FAIL, err.Error())
	}

	return count, customError.SUCCESS, nil
}

func GetCommentCount() (int64, int, error) {
	var count int64
	err := entity.Db.Find(&entity.Comment{}).Count(&count).Error
	if err != nil {
		return count, customError.COMMENT_LIST_FAIL, customError.GetError(customError.COMMENT_LIST_FAIL, err.Error())
	}

	return count, customError.SUCCESS, nil
}

func EnableComment(cid string) (int, error) {
	err := entity.Db.Where("cid =?", cid).Update("is_active", true).Error
	if err != nil {
		return customError.COMMENT_UPDATE_FAIL, customError.GetError(customError.COMMENT_UPDATE_FAIL, err.Error())
	}

	return customError.SUCCESS, nil
}
