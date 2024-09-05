package service

import (
	"slogv2/src/main/entity"
	"slogv2/src/main/utils/customError"
)

func GetRelationByAid(aid string) ([]entity.Relationship, int, error) {
	var relations []entity.Relationship
	err := entity.Db.Where("aid =?", aid).Find(&relations).Error
	if err != nil {
		return nil, customError.RELATION_NOT_FOUND, customError.GetError(customError.RELATION_NOT_FOUND, err.Error())
	}
	return relations, customError.SUCCESS, nil
}

func GetRelationByCid(cid string) ([]entity.Relationship, int, error) {
	var relations []entity.Relationship
	err := entity.Db.Where("cid=?", cid).Find(&relations).Error
	if err != nil {
		return nil, customError.RELATION_NOT_FOUND, customError.GetError(customError.RELATION_NOT_FOUND, err.Error())
	}
	return relations, customError.SUCCESS, nil
}

func CreateRelation(relation *entity.Relationship) (int, error) {
	err := entity.Db.Create(&relation).Error
	if err != nil {
		return customError.RELATION_CREATE_FAIL, customError.GetError(customError.RELATION_CREATE_FAIL, err.Error())
	}
	return customError.SUCCESS, nil
}

func DeleteRelation(relation *entity.Relationship) (int, error) {
	err := entity.Db.Delete(relation).Error
	if err != nil {
		return customError.RELATION_DELETE_FAIL, customError.GetError(customError.RELATION_CREATE_FAIL, err.Error())
	}
	return customError.SUCCESS, nil
}

func DeleteRelationByAid(aid string) (int, error) {
	err := entity.Db.Where("aid=?", aid).Delete(&entity.Relationship{}).Error
	if err != nil {
		return customError.RELATION_DELETE_FAIL, customError.GetError(customError.RELATION_CREATE_FAIL, err.Error())
	}
	return customError.SUCCESS, nil
}

func DeleteRelationByCid(cid string) (int, error) {
	err := entity.Db.Where("cid=?", cid).Delete(&entity.Relationship{}).Error
	if err != nil {
		return customError.RELATION_DELETE_FAIL, customError.GetError(customError.RELATION_CREATE_FAIL, err.Error())
	}
	return customError.SUCCESS, nil
}

func DeleteRelationList(relationList []entity.Relationship) (int, error) {
	err := entity.Db.Delete(&relationList).Error
	if err != nil {
		return customError.RELATION_DELETE_FAIL, customError.GetError(customError.RELATION_CREATE_FAIL, err.Error())
	}
	return customError.SUCCESS, nil
}
