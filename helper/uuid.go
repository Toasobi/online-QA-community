package helper

import uuid "github.com/satori/go.uuid"

//生成唯一标识
func GenerateUUID() string {
	return uuid.NewV4().String()
}
