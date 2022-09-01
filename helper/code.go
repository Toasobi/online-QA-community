package helper

import (
	"math/rand"
	"strconv"
	"time"
)

func GetCode() string {
	//随机数种子
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}

	return s
}
