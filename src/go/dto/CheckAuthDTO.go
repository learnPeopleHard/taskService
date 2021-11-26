package dto

import (
	"bytes"
	"encoding/gob"
)

type CheckAuthDTO struct{
	Token string
	FlagGetAndSet bool
}

func CheckEncode(data CheckAuthDTO) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func CheckDecode(b []byte) (CheckAuthDTO, error) {
	buf := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buf)
	var data CheckAuthDTO
	if err := decoder.Decode(&data); err != nil {
		return CheckAuthDTO{}, err
	}
	return data, nil
}
