package helper

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Identity           string `json:"identity"`
	Name               string `json:"name"`
	IsAdmin            int    `json:"is_admin"`
	jwt.StandardClaims        //必要
}

var myKey = []byte("online-QA-community") //密钥

//用户签发token
func GenerateToken(identity, name string, isAdmin int) (string, error) {
	userClaim := &UserClaims{
		Identity:       identity,
		Name:           name,
		IsAdmin:        isAdmin,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim) //这个token还不是我们最后想要的字符串token
	tokenString, err := token.SignedString(myKey)                 //注意要传byte数据
	if err != nil {
		return "", err
	}
	//fmt.Println(tokenString) //这样就生成了
	return tokenString, nil
}

//解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiR2V0IiwiaWRlbnRpdHkiOiJ1c2VyXzEifQ.OirWW9EXwML2aj6aqxvZETS3RtO7QVAvs0xI5eek2VI"
	userClaims := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaims, func(t *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !claims.Valid { //正常情况下打印
		//fmt.Println(userClaims)
		return nil, fmt.Errorf("Analyse TOken err:%v", err)
	}
	return userClaims, nil
}
