package respository

import (
	"context"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var ctx = context.Background()

//redis链接
var RDB *redis.Client

func InitRedis() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	db, _ := strconv.Atoi(viper.GetString("redis.database"))
	addr := strings.Join([]string{host, ":", port}, "")
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db, //use DB3
	})
	RDB = rdb
}
