package service

import (
	"log"
	"slogv2/src/main/entity"
	"slogv2/src/main/utils/customError"
	"strconv"
)

func CreateCategory(category *entity.Category) (int, error) {
	if category.CategoryName == "" {
		return customError.CATEGORY_CREATE_FAIL, customError.GetError(customError.CATEGORY_CREATE_FAIL, "分类名称不能为空")
	}

	var newCategory entity.Category
	newCategory.CategoryName = category.CategoryName
	newCategory.CategoryDesc = category.CategoryDesc
	newCategory.CategoryIcon = category.CategoryIcon
	newCategory.Count = 0
	newCategory.IsActive = true
	newCategory.ParentId = category.ParentId

	var c entity.Category
	var count int64
	err := entity.Db.Find(&c).Count(&count).Error
	if err != nil {
		return customError.CATEGORY_CREATE_FAIL, customError.GetError(customError.CATEGORY_CREATE_FAIL, err.Error())
	}
	if count == 0 {
		newCategory.CategoryId = 1
	} else {
		var _c entity.Category
		err = entity.Db.Last(&_c).Error
		if err != nil {
			return customError.CATEGORY_CREATE_FAIL, customError.GetError(customError.CATEGORY_CREATE_FAIL, err.Error())
		}
		newCategory.CategoryId = _c.CategoryId + 1
	}

	err = entity.Db.Create(&newCategory).Error
	if err != nil {
		return customError.CATEGORY_CREATE_FAIL, err
	}
	return customError.SUCCESS, nil
}

func GetCategoryList() ([]entity.Category, int, error) {
	var categoryList []entity.Category
	err := entity.Db.Find(&categoryList).Error
	if err != nil {
		return categoryList, customError.CATEGORY_LIST_FAIL, customError.GetError(customError.CATEGORY_LIST_FAIL, err.Error())
	}
	return categoryList, customError.SUCCESS, nil
}

func GetCategoryByCid(cid int) (entity.Category, int, error) {
	var category entity.Category
	err := entity.Db.Where("category_id =?", cid).Find(&category).Error
	if err != nil {
		return category, customError.CATEGORY_NOT_FOUND, customError.GetError(customError.CATEGORY_NOT_FOUND, err.Error())
	}
	return category, customError.SUCCESS, nil
}

func UpdateCategory(category *entity.Category) (int, error) {
	if category.CategoryName == "" {
		return customError.CATEGORY_UPDATE_FAIL, customError.GetError(customError.CATEGORY_UPDATE_FAIL, "分类名称不能为空")
	}

	err := entity.Db.Where("category_id =?", category.CategoryId).Updates(category).Error
	if err != nil {
		return customError.CATEGORY_UPDATE_FAIL, customError.GetError(customError.CATEGORY_UPDATE_FAIL, err.Error())
	}
	return customError.SUCCESS, nil
}

//func DeleteCategoryById(id string) (int, error) {
//	var category entity.Category
//	err := entity.Db.Where("category_id =?", id).Find(&category).Error
//	if err != nil {
//		return customError.CATEGORY_DELETE_FAIL, customError.GetError(customError.CATEGORY_DELETE_FAIL, err.Error())
//	}
//
//	//先寻找并删除子分类
//	var childCategoryList []entity.Category
//	err = entity.Db.Where("parent_id =?", id).Find(&childCategoryList).Error
//	if err != nil {
//		return customError.CATEGORY_DELETE_FAIL, customError.GetError(customError.CATEGORY_DELETE_FAIL, err.Error())
//	}
//
//	if len(childCategoryList) == 0 {
//		//无子分类时直接删除
//		//先删除分类
//		err = entity.Db.Where("parent_id =?", id).Delete(&entity.Category{}).Error
//		if err != nil {
//			return customError.CATEGORY_DELETE_FAIL, err
//		}
//
//		//再删除关系
//		status, err := DeleteRelationByCid(id)
//		if status!= customError.SUCCESS {
//			if err != nil {
//				return customError.CATEGORY_DELETE_FAIL, customError.GetError(customError.CATEGORY_DELETE_FAIL, err.Error())
//			}else{
//				return customError.CATEGORY_DELETE_FAIL, customError.GetError(customError.CATEGORY_DELETE_FAIL,
//					customError.GetMsg(status))
//			}
//		}else{
//			return customError.SUCCESS, nil
//		}
//
//	} else {
//		//有子分类时递归删除子分类
//		for _, childCategory := range childCategoryList {
//			status, err := DeleteCategoryById(strconv.Itoa(childCategory.CategoryId))
//
//			if err != nil {
//				return status, customError.GetError(status, err.Error())
//			}
//			if status == customError.SUCCESS {
//				continue
//			} else {
//				return status, customError.GetError(status, customError.GetMsg(status))
//			}
//		}
//	}
//
//	return customError.SUCCESS, nil
//}

//MarsCode优化加手动修改版

// DeleteCategoryById 递归删除分类
func DeleteCategoryById(id string) (int, error) {
	tx := entity.Db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("Recovered from panic:", r)
		}
	}()

	var category entity.Category
	err := tx.Where("category_id =?", id).Find(&category).Error
	if err != nil {
		tx.Rollback()
		return customError.CATEGORY_DELETE_FAIL, customError.GetError(customError.CATEGORY_DELETE_FAIL, err.Error())
	}

	// 先寻找并删除子分类
	var childCategoryList []entity.Category
	err = tx.Where("parent_id =?", id).Find(&childCategoryList).Error
	if err != nil {
		tx.Rollback()
		return customError.CATEGORY_DELETE_FAIL, customError.GetError(customError.CATEGORY_DELETE_FAIL, err.Error())
	}

	if len(childCategoryList) != 0 {
		// 有子分类时递归删除子分类
		childIds := make([]int, 0)
		for _, childCategory := range childCategoryList {
			childIds = append(childIds, childCategory.CategoryId)
		}

		err = tx.Where("category_id IN (?)", childIds).Delete(&entity.Category{}).Error
		if err != nil {
			tx.Rollback()
			return customError.CATEGORY_DELETE_FAIL, err
		}

		for _, childId := range childIds {
			status, err := DeleteCategoryById(strconv.Itoa(childId))
			if err != nil {
				tx.Rollback()
				return status, customError.GetError(status, err.Error())
			}
		}
	}

	// 无子分类时直接删除
	// 先删除分类
	err = tx.Where("parent_id =?", id).Delete(&entity.Category{}).Error
	if err != nil {
		tx.Rollback()
		return customError.CATEGORY_DELETE_FAIL, customError.GetError(customError.CATEGORY_DELETE_FAIL, err.Error())
	}

	// 再删除关系
	status, err := DeleteRelationByCid(id)
	if status != customError.SUCCESS {
		tx.Rollback()
		return status, err
	} else {
		tx.Commit()
		return customError.SUCCESS, nil
	}
}

func AddCategory(cid string, aid string) (int, error) {
	var rela entity.Relationship
	rela.Cid, _ = strconv.Atoi(cid)
	rela.Aid, _ = strconv.Atoi(aid)
	status, err := CreateRelation(&rela)
	if err != nil {
		return customError.RELATION_CREATE_FAIL, customError.GetError(customError.RELATION_CREATE_FAIL, err.Error())
	}
	return status, err
}

func RemoveCategory(cid string, aid string) (int, error) {
	var relation entity.Relationship
	relation.Cid, _ = strconv.Atoi(cid)
	relation.Aid, _ = strconv.Atoi(aid)
	status, err := DeleteRelation(&relation)
	if err != nil {
		return customError.RELATION_CREATE_FAIL, customError.GetError(customError.RELATION_CREATE_FAIL, err.Error())
	}
	return status, nil
}
