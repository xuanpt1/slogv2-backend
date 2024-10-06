package service

import (
	"slogv2/src/main/entity"
	"slogv2/src/main/utils"
	"slogv2/src/main/utils/customError"
	"strconv"
)

func CreateUser(user *entity.User) (int, error) {
	if user.Username == "" || user.Password == "" {
		return customError.USER_CREATE_FAIL, customError.GetError(customError.USER_CREATE_FAIL, "用户名或密码不能为空")
	}

	if user.Avatar == "" {
		user.Avatar = utils.TestDefaultImg
	}

	if user.Nickname == "" {
		user.Nickname = user.Username
	}

	//pwd, salt := scrypt.

	user.IsActive = true

	//手动实现主键UID自增
	var count int64
	err := entity.Db.Find(&entity.User{}).Count(&count).Error
	if err != nil {
		return customError.USER_CREATE_FAIL, customError.GetError(customError.USER_CREATE_FAIL, err.Error())
	}
	if count == 0 {
		user.Uid = "0"
	} else {
		var _user entity.User
		err := entity.Db.Last(&_user).Error
		if err != nil {
			return customError.USER_CREATE_FAIL, customError.GetError(customError.USER_CREATE_FAIL, err.Error())
		}
		lastUid, err := strconv.Atoi(_user.Uid)
		if err != nil {
			return customError.USER_CREATE_FAIL, customError.GetError(customError.USER_CREATE_FAIL, err.Error())
		}
		user.Uid = strconv.Itoa(lastUid + 1)
	}

	err = entity.Db.Create(user).Error
	if err != nil {
		return customError.USER_CREATE_FAIL, customError.GetError(customError.USER_CREATE_FAIL, err.Error())
	}

	return customError.SUCCESS, nil
}

func UpdateUser(user *entity.User) (int, error) {
	if user.Username == "" || user.Password == "" {
		return customError.USER_UPDATE_FAIL, customError.GetError(customError.USER_UPDATE_FAIL, "用户名或密码不能为空")
	}

	if user.Avatar == "" {
		user.Avatar = utils.TestDefaultImg
	}
	if user.Nickname == "" {
		user.Nickname = user.Username
	}

	err := entity.Db.Where("uid = ?", user.Uid).Save(user).Error
	if err != nil {
		return customError.USER_UPDATE_FAIL, customError.GetError(customError.USER_UPDATE_FAIL, err.Error())
	}
	return customError.SUCCESS, nil
}

func DeleteUser(user *entity.User) (int, error) {
	err := entity.Db.Where("uid =?", user.Uid).Delete(&entity.User{}).Error
	if err != nil {
		return customError.USER_DELETE_FAIL, customError.GetError(customError.USER_DELETE_FAIL, err.Error())
	}
	return customError.SUCCESS, nil
}

func GetUserByUid(uid string) (entity.User, int, error) {
	var user entity.User
	err := entity.Db.Where("uid =?", uid).First(&user).Error
	if err != nil {
		return user, customError.USER_NOT_FOUND, customError.GetError(customError.USER_NOT_FOUND, err.Error())
	}
	return user, customError.SUCCESS, nil
}

func GetUserByUsername(username string) (entity.User, int, error) {
	var user entity.User
	err := entity.Db.Where("username =?", username).First(&user).Error
	if err != nil {
		return user, customError.USER_NOT_FOUND, customError.GetError(customError.USER_NOT_FOUND, err.Error())
	}
	return user, customError.SUCCESS, nil
}

func GetUserList() ([]entity.User, int, error) {
	var userList []entity.User
	err := entity.Db.Find(&userList).Error
	if err != nil {
		return userList, customError.USER_LIST_FAIL, customError.GetError(customError.USER_LIST_FAIL, err.Error())
	}
	return userList, customError.SUCCESS, nil
}
