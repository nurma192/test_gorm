package database

import (
	"Go_rest_now_android/model"
	"errors"
)

func IsExistUserByLogin(login string) bool {
	var count = 0
	var user model.User

	GetDB().First(&user, model.User{Login: login}).Count(&count)

	return count == 1 && user.Id > 0
}

func Add(bean interface{}) error {
	if !GetDB().NewRecord(bean) {
		return errors.New("unable to create")
	}

	return GetDB().Create(bean).Error
}
