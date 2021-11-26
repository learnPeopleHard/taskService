package userinfodto

import (
	"bytes"
	"encoding/gob"
)

type UserInfoDTO struct{
	Id int64
	UserId int64
	UserName string
	Nickname string
	ProfilePicture string
}


func Decode(b []byte) (UserInfoDTO, error) {
	buf := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buf)
	var data UserInfoDTO
	if err := decoder.Decode(&data); err != nil {
		return UserInfoDTO{}, err
	}
	return data, nil
}