package service

import (
	"slogv2/src/main/entity"
	"slogv2/src/main/utils/customError"
)

func CreateOption(options *entity.Options) (int, error) {
	if options.Name == "" || options.Value == "" {
		return customError.OPTIONS_CREATE_FAIL, customError.GetError(customError.OPTIONS_CREATE_FAIL, "选项名或值不能为空")
	}

	err := entity.Db.Create(options).Error
	if err != nil {
		return customError.OPTIONS_CREATE_FAIL, customError.GetError(customError.OPTIONS_CREATE_FAIL, err.Error())
	}

	return customError.SUCCESS, nil
}

func UpdateOption(options *entity.Options) (int, error) {
	if options.Name == "" || options.Value == "" {
		return customError.OPTIONS_UPDATE_FAIL, customError.GetError(customError.OPTIONS_UPDATE_FAIL, "选项名或值不能为空")
	}

	err := entity.Db.Where("name = ?", options.Name).Updates(options).Error
	if err != nil {
		return customError.OPTIONS_UPDATE_FAIL, customError.GetError(customError.OPTIONS_UPDATE_FAIL, err.Error())
	}

	return customError.SUCCESS, nil
}

func DeleteOption(name string) (int, error) {
	err := entity.Db.Where("name = ?", name).Delete(&entity.Options{}).Error
	if err != nil {
		return customError.OPTIONS_DELETE_FAIL, customError.GetError(customError.OPTIONS_DELETE_FAIL, err.Error())
	}

	return customError.SUCCESS, nil
}

func GetOptionByKey(name string) (entity.Options, int, error) {
	var options entity.Options
	err := entity.Db.Where("name = ?", name).First(&options).Error
	if err != nil {
		return options, customError.OPTIONS_NOT_FOUND, customError.GetError(customError.OPTIONS_NOT_FOUND, err.Error())
	}

	return options, customError.SUCCESS, nil
}

func GetOptionList() ([]entity.Options, int, error) {
	var optionsList []entity.Options
	err := entity.Db.Find(&optionsList).Error
	if err != nil {
		return optionsList, customError.OPTIONS_LIST_FAIL, customError.GetError(customError.OPTIONS_LIST_FAIL, err.Error())
	}

	return optionsList, customError.SUCCESS, nil
}
