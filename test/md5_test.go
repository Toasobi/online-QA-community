package test

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(fmt.Sprintf("%x", md5.Sum([]byte("123456")))) //16进制返回
}
