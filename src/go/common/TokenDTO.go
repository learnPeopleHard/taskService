package common

import (
	"bytes"
	"encoding/gob"
)

type TokenDTO struct{
	UserId int64
	Token string
	LoginTime int64
	UserName string
}


func Decode(b []byte) (TokenDTO, error) {
	buf := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buf)
	var data TokenDTO
	if err := decoder.Decode(&data); err != nil {
		return TokenDTO{}, err
	}
	return data, nil
}