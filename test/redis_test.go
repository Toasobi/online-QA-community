package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

//redis链接
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       3, //use DB3
})

func TestRedisSet(t *testing.T) {
	// host := viper.GetString("redis.host")
	// port := viper.GetString("redis.port")
	// password := viper.GetString("redis.password")
	// db, _ := strconv.Atoi(viper.GetString("redis.database"))
	// addr := strings.Join([]string{host, ":", port}, "")

	rdb.Set(ctx, "name", "mmc", time.Second*300)
}

func TestRedisGet(t *testing.T) {
	v, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(v)

}
