package utils

import (
	"demo/common"
	"demo/model"
)

func IsTelephoneExist(telephone string) bool {
	var user model.User
	common.GetDB().Where("telephone=?", telephone).First(&user)
	if user.Telephone == telephone {
		return false
	}
	return true
}
